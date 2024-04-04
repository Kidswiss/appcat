package vshnkeycloakcontroller

import (
	"context"
	"fmt"
	"time"

	vshnv1 "github.com/vshn/appcat/v4/apis/vshn/v1"
	"github.com/vshn/appcat/v4/pkg/common/utils"
	"github.com/vshn/appcat/v4/pkg/sliexporter/probes"
	slireconciler "github.com/vshn/appcat/v4/pkg/sliexporter/sli_reconciler"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	ctrl "sigs.k8s.io/controller-runtime"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	vshnKeycloakServiceKey = "VSHNKeycloak"
)

type VSHNKeycloakReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	ProbeManager       probeManager
	StartupGracePeriod time.Duration
}

type probeManager interface {
	slireconciler.ProbeManager
}

//+kubebuilder:rbac:groups=vshn.appcat.vshn.io,resources=xvshnkeycloaks,verbs=get;list;watch
//+kubebuilder:rbac:groups=vshn.appcat.vshn.io,resources=xvshnkeycloaks/status,verbs=get

//+kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch
//+kubebuilder:rbac:groups="",resources=namespaces,verbs=get;list;watch

// Reconcile checks against the k8s internal service if Keycloak is avaialable.
// It uses an https connection to achieve this.
func (r *VSHNKeycloakReconciler) Reconcile(ctx context.Context, req ctrl.Request) (res ctrl.Result, err error) {
	l := log.FromContext(ctx).WithValues("namespace", req.Namespace, "instance", req.Name)
	l.Info("Reconciling XVSHNKeycloak")

	inst := &vshnv1.XVSHNKeycloak{}
	err = r.Get(ctx, req.NamespacedName, inst)

	reconciler := slireconciler.New(inst, l, r.ProbeManager, vshnKeycloakServiceKey, req.NamespacedName, err, r.StartupGracePeriod, r.getKeycloakProber)

	return reconciler.Reconcile(ctx)

}

func (r VSHNKeycloakReconciler) getKeycloakProber(ctx context.Context, obj slireconciler.Service) (prober probes.Prober, err error) {
	inst, ok := obj.(*vshnv1.XVSHNKeycloak)
	if !ok {
		return nil, fmt.Errorf("cannot start probe, object not a valid VSHNRedis")
	}

	credentials := corev1.Secret{}

	err = r.Get(ctx, types.NamespacedName{
		Name:      inst.GetWriteConnectionSecretToReference().Name,
		Namespace: inst.GetWriteConnectionSecretToReference().Namespace,
	}, &credentials)

	if err != nil {
		return nil, err
	}

	org := inst.GetLabels()[utils.OrgLabelName]
	sla := inst.Spec.Parameters.Service.ServiceLevel
	ha := inst.Spec.Parameters.Instances > 1

	host, ok := credentials.Data["KEYCLOAK_HOST"]
	if !ok {
		return nil, fmt.Errorf("secret does not contain Keycloak url")
	}

	url := "https://" + string(host) + ":8443/health"

	cacert, ok := credentials.Data["ca.crt"]
	if !ok {
		return nil, fmt.Errorf("secret does not contain ca certificate")
	}

	return probes.NewHTTP(url, true, cacert, vshnKeycloakServiceKey, inst.GetName(), inst.GetNamespace(), org, string(sla), ha), nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *VSHNKeycloakReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&vshnv1.XVSHNKeycloak{}).
		Complete(r)
}
