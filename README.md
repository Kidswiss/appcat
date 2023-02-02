# appcat-apiserve

## Generate Kubernetes code

If you make changes to the auto generated code you'll need to run code generation.
This can be done with make:

```bash
make generate
```

## Building

See `make help` for a list of build targets.

* `make build`: Build binary for linux/amd64
* `make build -e GOOS=darwin -e GOARCH=arm64`: Build binary for macos/arm64
* `make docker-build`: Build Docker image for local environment

## Local development environment

You can setup a [kind]-based local environment with

```bash
make local-install
```

See [docs](./docs/modules/ROOT/pages/tutorials/dev-environment.md) for more details on the local environment setup.

Please be aware that the productive deployment of the appcat-apiserver may run on a different Kubernetes distribution than [kind].

[kind]: https://kind.sigs.k8s.io/

## Debugging in IDE
To run the API server on your local machine you need to register the IDE running instance with kind cluster.
This can be achieved with the following guide.

The `externalName` needs to be changed to your specific host IP.
When running kind on Linux you can find it with `docker inspect`.
On some docker distributions the host IP is accessible via `host.docker.internal`.
For Lima distribution the host IP is accessible via `host.lima.internal`.

```bash
HOSTIP=$(docker inspect appcat-apiserver-v1.24.0-control-plane | jq '.[0].NetworkSettings.Networks.kind.Gateway')
# HOSTIP=host.docker.internal # On some docker distributions
# HOSTIP=host.lima.internal # On lima distributions

make kind-setup

kind get kubeconfig --name appcat-apiserver-v1.24.0  > ~/.kube/config 

kubectl apply -f https://raw.githubusercontent.com/crossplane/crossplane/master/cluster/crds/apiextensions.crossplane.io_compositions.yaml

cat <<EOF | sed -e "s/172.21.0.1/$HOSTIP/g" | kubectl apply -f -
apiVersion: apiregistration.k8s.io/v1
kind: APIService
metadata:
  name: v1.api.appcat.vshn.io
  labels:
    api: appcat
    apiserver: "true"
spec:
  version: v1
  group: api.appcat.vshn.io
  insecureSkipTLSVerify: true
  groupPriorityMinimum: 2000
  service:
    name: appcat
    namespace: default
    port: 9443
  versionPriority: 10
---
apiVersion: v1
kind: Service
metadata:
  name: appcat
  namespace: default
spec:
  ports:
  - port: 9443
    protocol: TCP
    targetPort: 9443
  type: ExternalName
  externalName: 172.21.0.1 # Change to host IP
EOF
```
