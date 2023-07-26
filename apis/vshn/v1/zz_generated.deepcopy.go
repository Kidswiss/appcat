//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	"github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1alpha1"
	apisv1 "github.com/vshn/appcat/apis/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8upBackupSpec) DeepCopyInto(out *K8upBackupSpec) {
	*out = *in
	out.Retention = in.Retention
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8upBackupSpec.
func (in *K8upBackupSpec) DeepCopy() *K8upBackupSpec {
	if in == nil {
		return nil
	}
	out := new(K8upBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8upRestoreSpec) DeepCopyInto(out *K8upRestoreSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8upRestoreSpec.
func (in *K8upRestoreSpec) DeepCopy() *K8upRestoreSpec {
	if in == nil {
		return nil
	}
	out := new(K8upRestoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8upRetentionPolicy) DeepCopyInto(out *K8upRetentionPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8upRetentionPolicy.
func (in *K8upRetentionPolicy) DeepCopy() *K8upRetentionPolicy {
	if in == nil {
		return nil
	}
	out := new(K8upRetentionPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNDBaaSMaintenanceScheduleSpec) DeepCopyInto(out *VSHNDBaaSMaintenanceScheduleSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNDBaaSMaintenanceScheduleSpec.
func (in *VSHNDBaaSMaintenanceScheduleSpec) DeepCopy() *VSHNDBaaSMaintenanceScheduleSpec {
	if in == nil {
		return nil
	}
	out := new(VSHNDBaaSMaintenanceScheduleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNDBaaSNetworkSpec) DeepCopyInto(out *VSHNDBaaSNetworkSpec) {
	*out = *in
	if in.IPFilter != nil {
		in, out := &in.IPFilter, &out.IPFilter
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNDBaaSNetworkSpec.
func (in *VSHNDBaaSNetworkSpec) DeepCopy() *VSHNDBaaSNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(VSHNDBaaSNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNDBaaSPostgresExtension) DeepCopyInto(out *VSHNDBaaSPostgresExtension) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNDBaaSPostgresExtension.
func (in *VSHNDBaaSPostgresExtension) DeepCopy() *VSHNDBaaSPostgresExtension {
	if in == nil {
		return nil
	}
	out := new(VSHNDBaaSPostgresExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNDBaaSSchedulingSpec) DeepCopyInto(out *VSHNDBaaSSchedulingSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNDBaaSSchedulingSpec.
func (in *VSHNDBaaSSchedulingSpec) DeepCopy() *VSHNDBaaSSchedulingSpec {
	if in == nil {
		return nil
	}
	out := new(VSHNDBaaSSchedulingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNDBaaSSizeRequestsSpec) DeepCopyInto(out *VSHNDBaaSSizeRequestsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNDBaaSSizeRequestsSpec.
func (in *VSHNDBaaSSizeRequestsSpec) DeepCopy() *VSHNDBaaSSizeRequestsSpec {
	if in == nil {
		return nil
	}
	out := new(VSHNDBaaSSizeRequestsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNDBaaSSizeSpec) DeepCopyInto(out *VSHNDBaaSSizeSpec) {
	*out = *in
	out.Requests = in.Requests
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNDBaaSSizeSpec.
func (in *VSHNDBaaSSizeSpec) DeepCopy() *VSHNDBaaSSizeSpec {
	if in == nil {
		return nil
	}
	out := new(VSHNDBaaSSizeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNPostgreSQL) DeepCopyInto(out *VSHNPostgreSQL) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNPostgreSQL.
func (in *VSHNPostgreSQL) DeepCopy() *VSHNPostgreSQL {
	if in == nil {
		return nil
	}
	out := new(VSHNPostgreSQL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VSHNPostgreSQL) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNPostgreSQLBackup) DeepCopyInto(out *VSHNPostgreSQLBackup) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNPostgreSQLBackup.
func (in *VSHNPostgreSQLBackup) DeepCopy() *VSHNPostgreSQLBackup {
	if in == nil {
		return nil
	}
	out := new(VSHNPostgreSQLBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNPostgreSQLEncryption) DeepCopyInto(out *VSHNPostgreSQLEncryption) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNPostgreSQLEncryption.
func (in *VSHNPostgreSQLEncryption) DeepCopy() *VSHNPostgreSQLEncryption {
	if in == nil {
		return nil
	}
	out := new(VSHNPostgreSQLEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNPostgreSQLList) DeepCopyInto(out *VSHNPostgreSQLList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VSHNPostgreSQL, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNPostgreSQLList.
func (in *VSHNPostgreSQLList) DeepCopy() *VSHNPostgreSQLList {
	if in == nil {
		return nil
	}
	out := new(VSHNPostgreSQLList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VSHNPostgreSQLList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNPostgreSQLMonitoring) DeepCopyInto(out *VSHNPostgreSQLMonitoring) {
	*out = *in
	if in.AlertmanagerConfigSpecTemplate != nil {
		in, out := &in.AlertmanagerConfigSpecTemplate, &out.AlertmanagerConfigSpecTemplate
		*out = new(v1alpha1.AlertmanagerConfigSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNPostgreSQLMonitoring.
func (in *VSHNPostgreSQLMonitoring) DeepCopy() *VSHNPostgreSQLMonitoring {
	if in == nil {
		return nil
	}
	out := new(VSHNPostgreSQLMonitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNPostgreSQLParameters) DeepCopyInto(out *VSHNPostgreSQLParameters) {
	*out = *in
	in.Service.DeepCopyInto(&out.Service)
	out.Maintenance = in.Maintenance
	out.Size = in.Size
	in.Scheduling.DeepCopyInto(&out.Scheduling)
	in.Network.DeepCopyInto(&out.Network)
	out.Backup = in.Backup
	out.Restore = in.Restore
	in.Monitoring.DeepCopyInto(&out.Monitoring)
	out.Encryption = in.Encryption
	out.UpdateStrategy = in.UpdateStrategy
	out.Replication = in.Replication
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNPostgreSQLParameters.
func (in *VSHNPostgreSQLParameters) DeepCopy() *VSHNPostgreSQLParameters {
	if in == nil {
		return nil
	}
	out := new(VSHNPostgreSQLParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNPostgreSQLReplicationStrategy) DeepCopyInto(out *VSHNPostgreSQLReplicationStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNPostgreSQLReplicationStrategy.
func (in *VSHNPostgreSQLReplicationStrategy) DeepCopy() *VSHNPostgreSQLReplicationStrategy {
	if in == nil {
		return nil
	}
	out := new(VSHNPostgreSQLReplicationStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNPostgreSQLRestore) DeepCopyInto(out *VSHNPostgreSQLRestore) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNPostgreSQLRestore.
func (in *VSHNPostgreSQLRestore) DeepCopy() *VSHNPostgreSQLRestore {
	if in == nil {
		return nil
	}
	out := new(VSHNPostgreSQLRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNPostgreSQLServiceSpec) DeepCopyInto(out *VSHNPostgreSQLServiceSpec) {
	*out = *in
	in.PostgreSQLSettings.DeepCopyInto(&out.PostgreSQLSettings)
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make([]VSHNDBaaSPostgresExtension, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNPostgreSQLServiceSpec.
func (in *VSHNPostgreSQLServiceSpec) DeepCopy() *VSHNPostgreSQLServiceSpec {
	if in == nil {
		return nil
	}
	out := new(VSHNPostgreSQLServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNPostgreSQLSpec) DeepCopyInto(out *VSHNPostgreSQLSpec) {
	*out = *in
	in.Parameters.DeepCopyInto(&out.Parameters)
	out.WriteConnectionSecretToRef = in.WriteConnectionSecretToRef
	out.ResourceRef = in.ResourceRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNPostgreSQLSpec.
func (in *VSHNPostgreSQLSpec) DeepCopy() *VSHNPostgreSQLSpec {
	if in == nil {
		return nil
	}
	out := new(VSHNPostgreSQLSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNPostgreSQLStatus) DeepCopyInto(out *VSHNPostgreSQLStatus) {
	*out = *in
	if in.PostgreSQLConditions != nil {
		in, out := &in.PostgreSQLConditions, &out.PostgreSQLConditions
		*out = make([]apisv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NamespaceConditions != nil {
		in, out := &in.NamespaceConditions, &out.NamespaceConditions
		*out = make([]apisv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProfileConditions != nil {
		in, out := &in.ProfileConditions, &out.ProfileConditions
		*out = make([]apisv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PGConfigConditions != nil {
		in, out := &in.PGConfigConditions, &out.PGConfigConditions
		*out = make([]apisv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PGClusterConditions != nil {
		in, out := &in.PGClusterConditions, &out.PGClusterConditions
		*out = make([]apisv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecretsConditions != nil {
		in, out := &in.SecretsConditions, &out.SecretsConditions
		*out = make([]apisv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ObjectBucketConditions != nil {
		in, out := &in.ObjectBucketConditions, &out.ObjectBucketConditions
		*out = make([]apisv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ObjectBackupConfigConditions != nil {
		in, out := &in.ObjectBackupConfigConditions, &out.ObjectBackupConfigConditions
		*out = make([]apisv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkPolicyConditions != nil {
		in, out := &in.NetworkPolicyConditions, &out.NetworkPolicyConditions
		*out = make([]apisv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LocalCAConditions != nil {
		in, out := &in.LocalCAConditions, &out.LocalCAConditions
		*out = make([]apisv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CertificateConditions != nil {
		in, out := &in.CertificateConditions, &out.CertificateConditions
		*out = make([]apisv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNPostgreSQLStatus.
func (in *VSHNPostgreSQLStatus) DeepCopy() *VSHNPostgreSQLStatus {
	if in == nil {
		return nil
	}
	out := new(VSHNPostgreSQLStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNPostgreSQLUpdateStrategy) DeepCopyInto(out *VSHNPostgreSQLUpdateStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNPostgreSQLUpdateStrategy.
func (in *VSHNPostgreSQLUpdateStrategy) DeepCopy() *VSHNPostgreSQLUpdateStrategy {
	if in == nil {
		return nil
	}
	out := new(VSHNPostgreSQLUpdateStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNRedis) DeepCopyInto(out *VSHNRedis) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNRedis.
func (in *VSHNRedis) DeepCopy() *VSHNRedis {
	if in == nil {
		return nil
	}
	out := new(VSHNRedis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VSHNRedis) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNRedisList) DeepCopyInto(out *VSHNRedisList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VSHNRedis, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNRedisList.
func (in *VSHNRedisList) DeepCopy() *VSHNRedisList {
	if in == nil {
		return nil
	}
	out := new(VSHNRedisList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VSHNRedisList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNRedisParameters) DeepCopyInto(out *VSHNRedisParameters) {
	*out = *in
	out.Service = in.Service
	out.Size = in.Size
	in.Scheduling.DeepCopyInto(&out.Scheduling)
	out.TLS = in.TLS
	out.Backup = in.Backup
	out.Restore = in.Restore
	out.Maintenance = in.Maintenance
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNRedisParameters.
func (in *VSHNRedisParameters) DeepCopy() *VSHNRedisParameters {
	if in == nil {
		return nil
	}
	out := new(VSHNRedisParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNRedisServiceSpec) DeepCopyInto(out *VSHNRedisServiceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNRedisServiceSpec.
func (in *VSHNRedisServiceSpec) DeepCopy() *VSHNRedisServiceSpec {
	if in == nil {
		return nil
	}
	out := new(VSHNRedisServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNRedisSizeSpec) DeepCopyInto(out *VSHNRedisSizeSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNRedisSizeSpec.
func (in *VSHNRedisSizeSpec) DeepCopy() *VSHNRedisSizeSpec {
	if in == nil {
		return nil
	}
	out := new(VSHNRedisSizeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNRedisSpec) DeepCopyInto(out *VSHNRedisSpec) {
	*out = *in
	in.Parameters.DeepCopyInto(&out.Parameters)
	out.WriteConnectionSecretToRef = in.WriteConnectionSecretToRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNRedisSpec.
func (in *VSHNRedisSpec) DeepCopy() *VSHNRedisSpec {
	if in == nil {
		return nil
	}
	out := new(VSHNRedisSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNRedisStatus) DeepCopyInto(out *VSHNRedisStatus) {
	*out = *in
	if in.NamespaceConditions != nil {
		in, out := &in.NamespaceConditions, &out.NamespaceConditions
		*out = make([]apisv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SelfSignedIssuerConditions != nil {
		in, out := &in.SelfSignedIssuerConditions, &out.SelfSignedIssuerConditions
		*out = make([]apisv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LocalCAConditions != nil {
		in, out := &in.LocalCAConditions, &out.LocalCAConditions
		*out = make([]apisv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CaCertificateConditions != nil {
		in, out := &in.CaCertificateConditions, &out.CaCertificateConditions
		*out = make([]apisv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServerCertificateConditions != nil {
		in, out := &in.ServerCertificateConditions, &out.ServerCertificateConditions
		*out = make([]apisv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClientCertificateConditions != nil {
		in, out := &in.ClientCertificateConditions, &out.ClientCertificateConditions
		*out = make([]apisv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNRedisStatus.
func (in *VSHNRedisStatus) DeepCopy() *VSHNRedisStatus {
	if in == nil {
		return nil
	}
	out := new(VSHNRedisStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSHNRedisTLSSpec) DeepCopyInto(out *VSHNRedisTLSSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSHNRedisTLSSpec.
func (in *VSHNRedisTLSSpec) DeepCopy() *VSHNRedisTLSSpec {
	if in == nil {
		return nil
	}
	out := new(VSHNRedisTLSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XVSHNPostgreSQL) DeepCopyInto(out *XVSHNPostgreSQL) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XVSHNPostgreSQL.
func (in *XVSHNPostgreSQL) DeepCopy() *XVSHNPostgreSQL {
	if in == nil {
		return nil
	}
	out := new(XVSHNPostgreSQL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XVSHNPostgreSQL) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XVSHNPostgreSQLList) DeepCopyInto(out *XVSHNPostgreSQLList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]XVSHNPostgreSQL, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XVSHNPostgreSQLList.
func (in *XVSHNPostgreSQLList) DeepCopy() *XVSHNPostgreSQLList {
	if in == nil {
		return nil
	}
	out := new(XVSHNPostgreSQLList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XVSHNPostgreSQLList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
