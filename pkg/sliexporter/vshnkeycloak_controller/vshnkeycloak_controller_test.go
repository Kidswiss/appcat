package vshnkeycloakcontroller

import (
	"context"
	"fmt"
	"testing"
	"time"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	vshnv1 "github.com/vshn/appcat/v4/apis/vshn/v1"
	"github.com/vshn/appcat/v4/pkg/sliexporter/probes"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

var _ probeManager = &fakeProbeManager{}

type fakeProbeManager struct {
	probers map[key]bool
}

func newFakeProbeManager() *fakeProbeManager {
	return &fakeProbeManager{
		probers: map[key]bool{},
	}
}

type key string

// StartProbe implements probeManager
func (m *fakeProbeManager) StartProbe(p probes.Prober) {
	m.probers[getFakeKey(p.GetInfo())] = true
}

// StopProbe implements probeManager
func (m *fakeProbeManager) StopProbe(p probes.ProbeInfo) {
	m.probers[getFakeKey(p)] = false
}

func getFakeKey(pi probes.ProbeInfo) key {
	return key(fmt.Sprintf("%s; %s", pi.Service, pi.Name))
}

func TestVSHNKeycloakReconciler_Reconcile(t *testing.T) {
	keycloak := newTestVSHNKeycloak("bar", "foo", "cred")

	r, manager, client := setupVSHNKeycloakTest(t,
		keycloak,
		newTestVSHNKeycloakCred("bar", "cred"))

	req := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Namespace: "bar",
			Name:      "foo",
		},
	}
	pi := probes.ProbeInfo{
		Service: "VSHNKeycloak",
		Name:    "foo",
	}

	_, err := r.Reconcile(context.TODO(), req)
	assert.NoError(t, err)
	assert.True(t, manager.probers[getFakeKey(pi)])

	require.NoError(t, client.Delete(context.TODO(), keycloak))
	_, err = r.Reconcile(context.TODO(), req)
	assert.NoError(t, err)
	assert.False(t, manager.probers[getFakeKey(pi)])
}

func setupVSHNKeycloakTest(t *testing.T, objs ...client.Object) (VSHNKeycloakReconciler, *fakeProbeManager, client.Client) {
	scheme := runtime.NewScheme()
	require.NoError(t, clientgoscheme.AddToScheme(scheme))
	require.NoError(t, vshnv1.AddToScheme(scheme))
	client := fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(objs...).
		Build()

	manager := newFakeProbeManager()
	r := VSHNKeycloakReconciler{
		Client:             client,
		Scheme:             scheme,
		StartupGracePeriod: 5 * time.Minute,
		ProbeManager:       manager,
	}

	return r, manager, client
}

func newTestVSHNKeycloak(namespace, name, cred string) *vshnv1.XVSHNKeycloak {
	return &vshnv1.XVSHNKeycloak{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: vshnv1.XVSHNKeycloakSpec{
			Parameters: vshnv1.VSHNKeycloakParameters{},
			ResourceSpec: xpv1.ResourceSpec{
				WriteConnectionSecretToReference: &xpv1.SecretReference{
					Name:      cred,
					Namespace: namespace,
				},
			},
		},
	}
}

func newTestVSHNKeycloakCred(namespace, name string) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: map[string][]byte{
			"KEYCLOAK_HOST": []byte("user"),
			"ca.crt": []byte(`-----BEGIN CERTIFICATE-----
MIIB3TCCAYOgAwIBAgIRAK2AWokJvb9o1OXYbU8CueYwCgYIKoZIzj0EAwIwTjEX
MBUGA1UEChMOdnNobi1hcHBjYXQtY2ExMzAxBgNVBAMTKmtleWNsb2FrLWFwcDEt
cHJvZC01aHA2NC1rZXljbG9ha3gtaHR0cC1jYTAeFw0yNDA0MDMxMjE3MjNaFw0z
NDA0MDExMjE3MjNaME4xFzAVBgNVBAoTDnZzaG4tYXBwY2F0LWNhMTMwMQYDVQQD
EyprZXljbG9hay1hcHAxLXByb2QtNWhwNjQta2V5Y2xvYWt4LWh0dHAtY2EwWTAT
BgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ9dP4bMvhr8ESprfX7Y6jlCUhOvlFeqd3S
v1sJuYCqYBTf87fg+pDOMAzsubdn8jyJUf65WwhcN2fOV4MesDZXo0IwQDAOBgNV
HQ8BAf8EBAMCAqQwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUvQClXJloKXVG
M11dLmi/FAOusZUwCgYIKoZIzj0EAwIDSAAwRQIhAOwQ5NAuEz7TQ5dEy41d7TFm
hbzWn3LKJJs7R13dKYJ8AiANYhF7QtPLyGxIkheciQsP+lQA+Yg4dfTfkgguaXHJ
rQ==
-----END CERTIFICATE-----`),
		},
	}
}
