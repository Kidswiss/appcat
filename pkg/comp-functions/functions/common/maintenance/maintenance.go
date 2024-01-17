package maintenance

import (
	"context"
	"fmt"
	"regexp"

	"github.com/blang/semver/v4"
	xkubev1 "github.com/crossplane-contrib/provider-kubernetes/apis/object/v1alpha1"
	xfnproto "github.com/crossplane/function-sdk-go/proto/v1beta1"
	vshnv1 "github.com/vshn/appcat/v4/apis/vshn/v1"
	"github.com/vshn/appcat/v4/pkg/comp-functions/runtime"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/utils/ptr"
	controllerruntime "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Maintenance contains data for maintenance k8s resource generation
type Maintenance struct {
	// instanceNamespace is the namespace where the service pod is running
	instanceNamespace string
	// mainRole is maintenance role name
	mainRole string
	// service is the service for which this maintenance is supposed to run. Ex:. postgresql
	service string
	// helmBasedService whether the maintenance is for a helm based service
	helmBasedService bool
	// resource object of this service
	resource client.Object
	// svc is is the ServiceRuntime
	svc *runtime.ServiceRuntime
	// schedule is the schedule spec of the resource
	schedule vshnv1.VSHNDBaaSMaintenanceScheduleSpec
	// policyRules are the permissions to be give to the maintenance role
	policyRules []rbacv1.PolicyRule
	// extraEnvs are extra environment variables to be added to the Cronjob
	extraEnvs []corev1.EnvVar
	// extraResources are extra resources to be added to the default list of resources created by this maintenance
	extraResources []ExtraResource
}

// ExtraResource is an extra resource to be added to the desired state of a Crossplane Function IO
type ExtraResource struct {
	Name     string
	Resource client.Object
	Refs     []xkubev1.Reference
}

var (
	maintServiceAccountName = "maintenanceserviceaccount"
	dayOfWeekMap            = map[string]int{
		"monday":    1,
		"tuesday":   2,
		"wednesday": 3,
		"thursday":  4,
		"friday":    5,
		"saturday":  6,
		"sunday":    0,
	}
)

// New creates a Maintenance object with required attributes
func New(r client.Object, svc *runtime.ServiceRuntime, schedule vshnv1.VSHNDBaaSMaintenanceScheduleSpec, instanceNamespace, service string) *Maintenance {
	return &Maintenance{
		instanceNamespace: instanceNamespace,
		service:           service,
		resource:          r,
		svc:               svc,
		schedule:          schedule,
	}
}

// WithPolicyRules sets the policy rules for the role
func (m *Maintenance) WithPolicyRules(policyRules []rbacv1.PolicyRule) *Maintenance {
	m.policyRules = policyRules
	return m
}

// WithHelmBasedService adds extra environment variables to the cron job
func (m *Maintenance) WithHelmBasedService() *Maintenance {
	m.helmBasedService = true
	return m
}

// WithExtraEnvs adds extra environment variables to the cron job
func (m *Maintenance) WithExtraEnvs(extraEnvs ...corev1.EnvVar) *Maintenance {
	m.extraEnvs = extraEnvs
	return m
}

// WithExtraResources adds extra resources to the desired composition function
func (m *Maintenance) WithExtraResources(extraResources ...ExtraResource) *Maintenance {
	m.extraResources = extraResources
	return m
}

// WithRole adds namespaced RBAC rules
func (m *Maintenance) WithRole(role string) *Maintenance {
	m.mainRole = role
	return m
}

// Run generates k8s resources for maintenance
func (m *Maintenance) Run(ctx context.Context) *xfnproto.Result {
	log := controllerruntime.LoggerFrom(ctx)

	log.Info("Adding maintenance cronjob to the instance")
	cron, err := m.parseCron()
	if err != nil {
		return runtime.NewFatalResult(err)
	}
	if cron == "" {
		log.Info("Maintenance schedule not yet populated")
		return runtime.NewNormalResult("Maintenance schedule not yet populated")
	}

	// Helm based services are having maintenance done in a control namespace therefore rbac rules are created
	// once in the component
	if !m.helmBasedService && m.mainRole != "" {
		err = m.createRBAC(ctx)
		if err != nil {
			return runtime.NewFatalResult(err)
		}
	}

	for _, extraR := range m.extraResources {
		err = m.svc.SetDesiredKubeObject(extraR.Resource, extraR.Name, runtime.KubeOptionAddRefs(extraR.Refs...))
		if err != nil {
			return runtime.NewFatalResult(err)
		}
	}

	err = m.createMaintenanceJob(ctx, cron)
	if err != nil {
		return runtime.NewFatalResult(err)
	}
	return nil
}

func (m *Maintenance) createRBAC(ctx context.Context) error {
	err := m.createMaintenanceServiceAccount(ctx)
	if err != nil {
		return fmt.Errorf("can't create maintenance serviceaccount: %v", err)
	}

	err = m.createMaintenanceRole(ctx)
	if err != nil {
		return fmt.Errorf("can't create maintenance role: %v", err)
	}

	err = m.createMaintenanceRolebinding(ctx)
	if err != nil {
		return fmt.Errorf("can't create maintenance rolebinding: %v", err)
	}
	return nil
}

func (m *Maintenance) createMaintenanceJob(ctx context.Context, cronSchedule string) error {
	imageTag := m.svc.Config.Data["imageTag"]
	if imageTag == "" {
		return fmt.Errorf("no imageTag field in composition function configuration")
	}

	sa := maintServiceAccountName
	jobNamespace := m.instanceNamespace
	jobName := "maintenancejob"

	// For helm based services create the job in the control namespace
	if m.helmBasedService {
		jobName = m.resource.GetName() + "-maintenancejob"
		jobNamespace = m.svc.Config.Data["controlNamespace"]
		if jobNamespace == "" {
			return fmt.Errorf("no controlNamespace field in composition function configuration")
		}
		sa = m.svc.Config.Data["maintenanceSA"]
		if sa == "" {
			return fmt.Errorf("no maintenanceSA field in composition function configuration")
		}
	}

	envVars := []corev1.EnvVar{
		{
			Name:  "INSTANCE_NAMESPACE",
			Value: m.instanceNamespace,
		},
		{
			Name:  "CLAIM_NAME",
			Value: m.resource.GetLabels()["crossplane.io/claim-name"],
		},
		{
			Name:  "CLAIM_NAMESPACE",
			Value: m.resource.GetLabels()["crossplane.io/claim-namespace"],
		},
	}

	job := &batchv1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      jobName,
			Namespace: jobNamespace,
		},
		Spec: batchv1.CronJobSpec{
			Schedule:                   cronSchedule,
			SuccessfulJobsHistoryLimit: ptr.To(int32(0)),
			JobTemplate: batchv1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							ServiceAccountName: sa,
							RestartPolicy:      corev1.RestartPolicyNever,
							Containers: []corev1.Container{
								{
									Name:  "maintenancejob",
									Image: "ghcr.io/vshn/appcat:" + imageTag,
									Env:   append(envVars, m.extraEnvs...),
									Args: []string{
										"maintenance",
										"--service",
										m.service,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return m.svc.SetDesiredKubeObject(job, m.resource.GetName()+"-maintenancejob")
}

func (m *Maintenance) createMaintenanceRolebinding(ctx context.Context) error {
	roleBinding := &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.mainRole,
			Namespace: m.instanceNamespace,
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
			Name:     m.mainRole,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind: "ServiceAccount",
				Name: maintServiceAccountName,
			},
		},
	}

	return m.svc.SetDesiredKubeObject(roleBinding, m.resource.GetName()+"-maintenance-rolebinding")
}

func (m *Maintenance) createMaintenanceRole(ctx context.Context) error {
	role := &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.mainRole,
			Namespace: m.instanceNamespace,
		},
		Rules: m.policyRules,
	}

	return m.svc.SetDesiredKubeObject(role, m.resource.GetName()+"-maintenance-role")
}

func (m *Maintenance) createMaintenanceServiceAccount(ctx context.Context) error {
	sa := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      maintServiceAccountName,
			Namespace: m.instanceNamespace,
		},
	}

	return m.svc.SetDesiredKubeObject(sa, m.resource.GetName()+"-maintenance-serviceaccount")
}

func (m *Maintenance) parseCron() (string, error) {

	if m.schedule.DayOfWeek == "" || m.schedule.TimeOfDay == "" {
		return "", nil
	}

	cronDayOfWeek := dayOfWeekMap[m.schedule.DayOfWeek]

	r := regexp.MustCompile(`(\d+):(\d+):.*`)
	timeSlice := r.FindStringSubmatch(m.schedule.TimeOfDay)

	if len(timeSlice) == 0 {
		return "", fmt.Errorf("not a valid time string %s", m.schedule.TimeOfDay)
	}

	return fmt.Sprintf("%s %s * * %d", timeSlice[2], timeSlice[1], cronDayOfWeek), nil
}

// SetReleaseVersion sets the version from the claim if it's a new instance otherwise it is managed by maintenance function
func SetReleaseVersion(ctx context.Context, version string, desiredValues map[string]interface{}, observedValues map[string]interface{}, fields []string) error {
	l := controllerruntime.LoggerFrom(ctx)

	tag, _, err := unstructured.NestedString(observedValues, fields...)
	if err != nil {
		return fmt.Errorf("cannot get image tag from values in release: %v", err)
	}

	desiredVersion, err := semver.ParseTolerant(version)
	if err != nil {
		l.Info("failed to parse desired service version", "version", version)
		return fmt.Errorf("invalid service version %q", version)
	}

	observedVersion, err := semver.ParseTolerant(tag)
	if err != nil {
		l.Info("failed to parse observed service version", "version", tag)
		// If the observed version is not parsable, e.g. if it's empty, update to the desired version
		return unstructured.SetNestedField(desiredValues, version, fields...)
	}

	if observedVersion.GTE(desiredVersion) {
		// In case the overved tag is valid and greater than the desired version, keep the observed version
		return unstructured.SetNestedField(desiredValues, tag, fields...)
	}
	// In case the observed tag is smaller than the desired version,  then set the version from the claim
	return unstructured.SetNestedField(desiredValues, version, fields...)
}
