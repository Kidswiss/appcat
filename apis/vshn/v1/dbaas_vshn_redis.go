package v1

import (
	"fmt"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	v1 "github.com/vshn/appcat/v4/apis/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Workaround to make nested defaulting work.
// kubebuilder is unable to set a {} default
//go:generate yq -i e ../../generated/vshn.appcat.vshn.io_vshnredis.yaml --expression "with(.spec.versions[]; .schema.openAPIV3Schema.properties.spec.properties.parameters.default={})"
//go:generate yq -i e ../../generated/vshn.appcat.vshn.io_vshnredis.yaml --expression "with(.spec.versions[]; .schema.openAPIV3Schema.properties.spec.properties.parameters.properties.size.default={})"
//go:generate yq -i e ../../generated/vshn.appcat.vshn.io_vshnredis.yaml --expression "with(.spec.versions[]; .schema.openAPIV3Schema.properties.spec.properties.parameters.properties.service.default={})"
//go:generate yq -i e ../../generated/vshn.appcat.vshn.io_vshnredis.yaml --expression "with(.spec.versions[]; .schema.openAPIV3Schema.properties.spec.properties.parameters.properties.tls.default={})"
//go:generate yq -i e ../../generated/vshn.appcat.vshn.io_vshnredis.yaml --expression "with(.spec.versions[]; .schema.openAPIV3Schema.properties.spec.properties.parameters.properties.backup.default={})"

// +kubebuilder:object:root=true

// VSHNRedis is the API for creating Redis clusters.
type VSHNRedis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of a VSHNRedis.
	Spec VSHNRedisSpec `json:"spec"`

	// Status reflects the observed state of a VSHNRedis.
	Status VSHNRedisStatus `json:"status,omitempty"`
}

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true
type VSHNRedisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []VSHNRedis `json:"items,omitempty"`
}

// VSHNRedisSpec defines the desired state of a VSHNRedis.
type VSHNRedisSpec struct {
	// Parameters are the configurable fields of a VSHNRedis.
	Parameters VSHNRedisParameters `json:"parameters,omitempty"`

	// WriteConnectionSecretToRef references a secret to which the connection details will be written.
	WriteConnectionSecretToRef v1.LocalObjectReference `json:"writeConnectionSecretToRef,omitempty"`
}

// VSHNRedisParameters are the configurable fields of a VSHNRedis.
type VSHNRedisParameters struct {
	// Service contains Redis DBaaS specific properties
	Service VSHNRedisServiceSpec `json:"service,omitempty"`

	// Size contains settings to control the sizing of a service.
	Size VSHNRedisSizeSpec `json:"size,omitempty"`

	// Scheduling contains settings to control the scheduling of an instance.
	Scheduling VSHNDBaaSSchedulingSpec `json:"scheduling,omitempty"`

	// TLS contains settings to control tls traffic of a service.
	TLS VSHNRedisTLSSpec `json:"tls,omitempty"`

	// Backup contains settings to control how the instance should get backed up.
	Backup K8upBackupSpec `json:"backup,omitempty"`

	// Restore contains settings to control the restore of an instance.
	Restore K8upRestoreSpec `json:"restore,omitempty"`

	// Maintenance contains settings to control the maintenance of an instance.
	Maintenance VSHNDBaaSMaintenanceScheduleSpec `json:"maintenance,omitempty"`

	// Monitoring contains settings to control monitoring.
	Monitoring VSHNMonitoring `json:"monitoring,omitempty"`

	// Security defines the security of a service
	Security Security `json:"security,omitempty"`
}

// VSHNRedisServiceSpec contains Redis DBaaS specific properties
type VSHNRedisServiceSpec struct {
	// +kubebuilder:validation:Enum="6.2";"7.0"
	// +kubebuilder:default="7.0"

	// Version contains supported version of Redis.
	// Multiple versions are supported. The latest version "7.0" is the default version.
	Version string `json:"version,omitempty"`

	// RedisSettings contains additional Redis settings.
	RedisSettings string `json:"redisSettings,omitempty"`

	// +kubebuilder:validation:Enum="besteffort";"guaranteed"
	// +kubebuilder:default="besteffort"

	// ServiceLevel defines the service level of this service. Either Best Effort or Guaranteed Availability is allowed.
	ServiceLevel VSHNDBaaSServiceLevel `json:"serviceLevel,omitempty"`
}

// VSHNRedisSizeSpec contains settings to control the sizing of a service.
type VSHNRedisSizeSpec struct {

	// CPURequests defines the requests amount of Kubernetes CPUs for an instance.
	CPURequests string `json:"cpuRequests,omitempty"`

	// CPULimits defines the limits amount of Kubernetes CPUs for an instance.
	CPULimits string `json:"cpuLimits,omitempty"`

	// MemoryRequests defines the requests amount of memory in units of bytes for an instance.
	MemoryRequests string `json:"memoryRequests,omitempty"`

	// MemoryLimits defines the limits amount of memory in units of bytes for an instance.
	MemoryLimits string `json:"memoryLimits,omitempty"`

	// Disk defines the amount of disk space for an instance.
	Disk string `json:"disk,omitempty"`

	// Plan is the name of the resource plan that defines the compute resources.
	Plan string `json:"plan,omitempty"`
}

// VSHNRedisTLSSpec contains settings to control tls traffic of a service.
type VSHNRedisTLSSpec struct {
	// +kubebuilder:default=true

	// TLSEnabled enables TLS traffic for the service
	TLSEnabled bool `json:"enabled,omitempty"`

	// +kubebuilder:default=true
	// TLSAuthClients enables client authentication requirement
	TLSAuthClients bool `json:"authClients,omitempty"`
}

// VSHNRedisStatus reflects the observed state of a VSHNRedis.
type VSHNRedisStatus struct {
	// RedisConditions contains the status conditions of the backing object.
	NamespaceConditions         []v1.Condition `json:"namespaceConditions,omitempty"`
	SelfSignedIssuerConditions  []v1.Condition `json:"selfSignedIssuerConditions,omitempty"`
	LocalCAConditions           []v1.Condition `json:"localCAConditions,omitempty"`
	CaCertificateConditions     []v1.Condition `json:"caCertificateConditions,omitempty"`
	ServerCertificateConditions []v1.Condition `json:"serverCertificateConditions,omitempty"`
	ClientCertificateConditions []v1.Condition `json:"clientCertificateConditions,omitempty"`
	// InstanceNamespace contains the name of the namespace where the instance resides
	InstanceNamespace string `json:"instanceNamespace,omitempty"`
	// Schedules keeps track of random generated schedules, is overwriten by
	// schedules set in the service's spec.
	Schedules VSHNScheduleStatus `json:"schedules,omitempty"`
}

func (v *VSHNRedis) GetClaimNamespace() string {
	return v.GetLabels()["crossplane.io/claim-namespace"]
}

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true

// XVSHNRedis represents the internal composite of this claim
type XVSHNRedis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   XVSHNRedisSpec   `json:"spec"`
	Status XVSHNRedisStatus `json:"status,omitempty"`
}

// XVSHNRedisSpec defines the desired state of a VSHNRedis.
type XVSHNRedisSpec struct {
	// Parameters are the configurable fields of a VSHNRedis.
	Parameters VSHNRedisParameters `json:"parameters,omitempty"`

	xpv1.ResourceSpec `json:",inline"`
}

type XVSHNRedisStatus struct {
	VSHNRedisStatus     `json:",inline"`
	xpv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true

// XVSHNRedisList represents a list of composites
type XVSHNRedisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []XVSHNRedis `json:"items"`
}

func (v *VSHNRedis) GetVSHNMonitoring() VSHNMonitoring {
	return v.Spec.Parameters.Monitoring
}

func (v *VSHNRedis) GetInstanceNamespace() string {
	return fmt.Sprintf("vshn-redis-%s", v.GetName())
	// GetMaintenanceDayOfWeek returns the currently set day of week
}

func (v *VSHNRedis) SetInstanceNamespaceStatus() {
	v.Status.InstanceNamespace = v.GetInstanceNamespace()
}

func (v *VSHNRedis) GetMaintenanceDayOfWeek() string {
	if v.Spec.Parameters.Maintenance.DayOfWeek != "" {
		return v.Spec.Parameters.Maintenance.DayOfWeek
	}
	return v.Status.Schedules.Maintenance.DayOfWeek
}

// GetMaintenanceTimeOfDay returns the currently set time of day
func (v *VSHNRedis) GetMaintenanceTimeOfDay() string {
	if v.Spec.Parameters.Maintenance.TimeOfDay != "" {
		return v.Spec.Parameters.Maintenance.TimeOfDay
	}
	return v.Status.Schedules.Maintenance.TimeOfDay
}

// SetMaintenanceDayOfWeek sets the day of week to the given value
func (v *VSHNRedis) SetMaintenanceDayOfWeek(dow string) {
	v.Status.Schedules.Maintenance.DayOfWeek = dow
}

// SetMaintenanceTimeOfDay sets the time of day to the given value
func (v *VSHNRedis) SetMaintenanceTimeOfDay(tod string) {
	v.Status.Schedules.Maintenance.TimeOfDay = tod
}

// GetBackupSchedule returns the current backup schedule
func (v *VSHNRedis) GetBackupSchedule() string {
	if v.Spec.Parameters.Backup.Schedule != "" {
		return v.Spec.Parameters.Backup.Schedule
	}
	return v.Status.Schedules.Backup
}

// SetBackupSchedule overwrites the current backup schedule
func (v *VSHNRedis) SetBackupSchedule(schedule string) {
	v.Status.Schedules.Backup = schedule
}

// GetBackupRetention returns the retention definition for this backup.
func (v *VSHNRedis) GetBackupRetention() K8upRetentionPolicy {
	return v.Spec.Parameters.Backup.Retention
}

// GetServiceName returns the name of this service
func (v *VSHNRedis) GetServiceName() string {
	return "redis"
}

// GetFullMaintenanceSchedule returns
func (v *VSHNRedis) GetFullMaintenanceSchedule() VSHNDBaaSMaintenanceScheduleSpec {
	schedule := v.Spec.Parameters.Maintenance
	schedule.DayOfWeek = v.GetMaintenanceDayOfWeek()
	schedule.TimeOfDay = v.GetMaintenanceTimeOfDay()
	return schedule
}

// GetAllowedNamespaces returns allowed namespaces to access this service
func (v *VSHNRedis) GetAllowedNamespaces() []string {
	return v.Spec.Parameters.Security.AllowedNamespaces
}
