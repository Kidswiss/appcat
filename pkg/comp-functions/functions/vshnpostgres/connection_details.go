package vshnpostgres

import (
	"context"
	"fmt"

	// "github.com/crossplane/crossplane/apis/apiextensions/fn/io/v1alpha1"
	xkubev1 "github.com/crossplane-contrib/provider-kubernetes/apis/object/v1alpha2"
	commonv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/function-sdk-go/proto/v1beta1"
	vshnv1 "github.com/vshn/appcat/v4/apis/vshn/v1"
	"github.com/vshn/appcat/v4/pkg/comp-functions/runtime"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	controllerruntime "sigs.k8s.io/controller-runtime"
)

const (
	// PostgresqlHost is env variable in the connection secret
	PostgresqlHost = "POSTGRESQL_HOST"
	// PostgresqlUser is env variable in the connection secret
	PostgresqlUser = "POSTGRESQL_USER"
	// PostgresqlPassword is env variable in the connection secret
	PostgresqlPassword = "POSTGRESQL_PASSWORD"
	// PostgresqlPort is env variable in the connection secret
	PostgresqlPort = "POSTGRESQL_PORT"
	// PostgresqlDb is env variable in the connection secret
	PostgresqlDb = "POSTGRESQL_DB"
	// PostgresqlURL is env variable in the connection secret
	PostgresqlURL = "POSTGRESQL_URL"
	defaultUser   = "postgres"
	defaultPort   = "5432"
	defaultDB     = "postgres"
)

// AddConnectionDetails changes the desired state of a FunctionIO
func AddConnectionDetails(ctx context.Context, svc *runtime.ServiceRuntime) *v1beta1.Result {
	log := controllerruntime.LoggerFrom(ctx)

	comp := &vshnv1.VSHNPostgreSQL{}
	err := svc.GetObservedComposite(comp)
	if err != nil {
		return runtime.NewFatalResult(fmt.Errorf("cannot get composite: %w", err))
	}

	log.Info("Making sure the cluster exposed connection details")
	obj := &xkubev1.Object{}
	err = svc.GetObservedComposedResource(obj, "cluster")
	if err != nil {
		return runtime.NewWarningResult(fmt.Sprintf("cannot get the sgcluster object: %s", err))
	}

	err = addConnectionDetailsToObject(obj, comp, svc)
	if err != nil {
		return runtime.NewWarningResult(fmt.Sprintf("cannot expose connection details on cluster: %s", err))
	}

	log.Info("Creating connection details")
	cd, err := svc.GetObservedComposedResourceConnectionDetails("cluster")
	if err != nil {
		return runtime.NewWarningResult(fmt.Sprintf("cannot get credentials from cluster object: %s", err))
	}

	if len(cd) == 0 {
		return runtime.NewWarningResult("no connection details yet on cluster")
	}

	host := fmt.Sprintf("%s.vshn-postgresql-%s.svc.cluster.local", comp.GetName(), comp.GetName())

	url := getPostgresURL(cd, host)

	svc.SetConnectionDetail(PostgresqlURL, []byte(url))
	svc.SetConnectionDetail(PostgresqlDb, []byte(defaultDB))
	svc.SetConnectionDetail(PostgresqlPort, []byte(defaultPort))
	svc.SetConnectionDetail(PostgresqlPassword, cd[PostgresqlPassword])
	svc.SetConnectionDetail(PostgresqlUser, []byte(defaultUser))
	svc.SetConnectionDetail(PostgresqlHost, []byte(host))
	err = svc.AddObservedConnectionDetails("cluster")
	if err != nil {
		return runtime.NewWarningResult(fmt.Sprintf("cannot add connection details to composite: %s", err))
	}

	return nil
}

func getPostgresURL(s map[string][]byte, host string) string {
	return getPostgresURLCustomUser(s, host, defaultUser)
}

func getPostgresURLCustomUser(s map[string][]byte, host, userName string) string {
	pwd := string(s[PostgresqlPassword])

	// The values are still missing, wait for the next reconciliation
	if pwd == "" {
		return ""
	}

	return "postgres://" + userName + ":" + pwd + "@" + host + ":" + defaultPort + "/" + defaultDB
}

func addConnectionDetailsToObject(obj *xkubev1.Object, comp *vshnv1.VSHNPostgreSQL, svc *runtime.ServiceRuntime) error {
	certSecretName := "tls-certificate"

	obj.Spec.ConnectionDetails = []xkubev1.ConnectionDetail{
		{
			ToConnectionSecretKey: PostgresqlPassword,
			ObjectReference: corev1.ObjectReference{
				APIVersion: "v1",
				Kind:       "Secret",
				Namespace:  comp.GetInstanceNamespace(),
				Name:       comp.GetName(),
				FieldPath:  "data.superuser-password",
			},
		},
		{
			ToConnectionSecretKey: "ca.crt",
			ObjectReference: corev1.ObjectReference{
				APIVersion: "v1",
				Kind:       "Secret",
				Namespace:  comp.GetInstanceNamespace(),
				Name:       certSecretName,
				FieldPath:  "data[ca.crt]",
			},
		},
		{
			ToConnectionSecretKey: "tls.crt",
			ObjectReference: corev1.ObjectReference{
				APIVersion: "v1",
				Kind:       "Secret",
				Namespace:  comp.GetInstanceNamespace(),
				Name:       certSecretName,
				FieldPath:  "data[tls.crt]",
			},
		},
		{
			ToConnectionSecretKey: "tls.key",
			ObjectReference: corev1.ObjectReference{
				APIVersion: "v1",
				Kind:       "Secret",
				Namespace:  comp.GetInstanceNamespace(),
				Name:       certSecretName,
				FieldPath:  "data[tls.key]",
			},
		},
	}

	obj.Spec.WriteConnectionSecretToReference = &commonv1.SecretReference{
		Name:      comp.GetName() + "-connection",
		Namespace: svc.Config.Data["crossplaneNamespace"],
	}

	err := svc.SetDesiredComposedResourceWithName(obj, "cluster")
	if err != nil {
		return err
	}

	// TODO: should probably go somewhere else...
	cd, err := svc.GetObservedComposedResourceConnectionDetails("pg-bucket")
	if err != nil {
		return err
	}

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pgbucket-" + comp.GetName(),
			Namespace: comp.GetInstanceNamespace(),
		},
		Data: cd,
	}

	return svc.SetDesiredKubeObject(secret, comp.GetName()+"-bucket-credentials")
}
