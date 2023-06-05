
# Image URL to use all building/pushing image targets
IMG_TAG ?= latest
GHCR_IMG ?= ghcr.io/vshn/appcat:$(IMG_TAG)
DOCKER_CMD ?= docker

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

OS := $(shell uname)
ifeq ($(OS), Darwin)
	sed ?= gsed
else
	sed ?= sed
endif

# For alpine image it is required the following env before building the application
DOCKER_IMAGE_GOOS = linux
DOCKER_IMAGE_GOARCH = amd64

PROJECT_ROOT_DIR = .
PROJECT_NAME ?= appcat
PROJECT_OWNER ?= vshn

PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
BIN_FILENAME ?= $(PROJECT_DIR)/appcat

## Stackgres CRDs
STACKGRES_VERSION ?= 1.4.3
STACKGRES_CRD_URL ?= https://gitlab.com/ongresinc/stackgres/-/raw/${STACKGRES_VERSION}/stackgres-k8s/src/common/src/main/resources/crds

## BUILD:go
go_bin ?= $(PWD)/.work/bin
$(go_bin):
	@mkdir -p $@

uname_s := $(shell uname -s)
ifeq ($(uname_s),Linux)
	distr_protoc := linux-x86_64
else
	distr_protoc := osx-universal_binary
endif

protoc_bin = $(go_bin)/protoc
$(protoc_bin): export GOBIN = $(go_bin)
$(protoc_bin): | $(go_bin)
	@echo "installing protocol buffers with dependencies"
	@git clone -q --depth 1 https://github.com/kubernetes/kubernetes.git .work/kubernetes
	@go install github.com/gogo/protobuf/protoc-gen-gogo@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	@wget -q -O $(go_bin)/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v22.1/protoc-22.1-$(distr_protoc).zip
	@unzip $(go_bin)/protoc.zip -d .work
	@rm $(go_bin)/protoc.zip

-include docs/antora-preview.mk docs/antora-build.mk

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: generate
generate: export PATH := $(go_bin):$(PATH)
generate: $(protoc_bin) generate-stackgres-crds ## Generate code with controller-gen and protobuf.
	go version
	rm -rf apis/generated
	go run sigs.k8s.io/controller-tools/cmd/controller-gen paths=./apis/... object crd:crdVersions=v1,allowDangerousTypes=true output:artifacts:config=./apis/generated
	go generate ./...
	# Because yaml is such a fun and easy specification, we need to hack some things here.
	# Depending on the yaml parser implementation the equal sign (=) has special meaning, or not...
	# So we make it explicitly a string.
	$(sed) -i ':a;N;$$!ba;s/- =\n/- "="\n/g' apis/generated/vshn.appcat.vshn.io_vshnpostgresqls.yaml
	rm -rf crds && cp -r apis/generated crds
	go run sigs.k8s.io/controller-tools/cmd/controller-gen rbac:roleName=appcat paths="{./apis/...,./pkg/apiserver/...}" output:artifacts:config=config/apiserver
	go run k8s.io/code-generator/cmd/go-to-protobuf \
		--packages=github.com/vshn/appcat/apis/appcat/v1 \
		--output-base=./.work/tmp \
		--go-header-file=./pkg/apiserver/hack/boilerplate.txt  \
        --apimachinery-packages='-k8s.io/apimachinery/pkg/util/intstr,-k8s.io/apimachinery/pkg/api/resource,-k8s.io/apimachinery/pkg/runtime/schema,-k8s.io/apimachinery/pkg/runtime,-k8s.io/apimachinery/pkg/apis/meta/v1,-k8s.io/apimachinery/pkg/apis/meta/v1beta1,-k8s.io/api/core/v1,-k8s.io/api/rbac/v1' \
        --proto-import=./.work/kubernetes/vendor/ && \
    	mv ./.work/tmp/github.com/vshn/appcat/apis/appcat/v1/generated.pb.go ./apis/appcat/v1/ && \
    	rm -rf ./.work/tmp

.PHONY: generate-stackgres-crds
generate-stackgres-crds:
	curl ${STACKGRES_CRD_URL}/SGDbOps.yaml?inline=false -o apis/stackgres/v1/sgdbops_crd.yaml
	yq -i e apis/stackgres/v1/sgdbops.yaml --expression ".components.schemas.SGDbOpsSpec=load(\"apis/stackgres/v1/sgdbops_crd.yaml\").spec.versions[0].schema.openAPIV3Schema.properties.spec"
	yq -i e apis/stackgres/v1/sgdbops.yaml --expression ".components.schemas.SGDbOpsStatus=load(\"apis/stackgres/v1/sgdbops_crd.yaml\").spec.versions[0].schema.openAPIV3Schema.properties.status"
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=v1 -generate=types -o apis/stackgres/v1/sgdbops.gen.go apis/stackgres/v1/sgdbops.yaml
	perl -i -0pe 's/\*struct\s\{\n\s\sAdditionalProperties\smap\[string\]string\s`json:"-"`\n\s}/map\[string\]string/gms' apis/stackgres/v1/sgdbops.gen.go

	curl ${STACKGRES_CRD_URL}/SGCluster.yaml?inline=false -o apis/stackgres/v1/sgcluster_crd.yaml
	yq -i e apis/stackgres/v1/sgcluster.yaml --expression ".components.schemas.SGClusterSpec=load(\"apis/stackgres/v1/sgcluster_crd.yaml\").spec.versions[0].schema.openAPIV3Schema.properties.spec"
	yq -i e apis/stackgres/v1/sgcluster.yaml --expression ".components.schemas.SGClusterStatus=load(\"apis/stackgres/v1/sgcluster_crd.yaml\").spec.versions[0].schema.openAPIV3Schema.properties.status"
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=v1 -generate=types -o apis/stackgres/v1/sgcluster.gen.go apis/stackgres/v1/sgcluster.yaml
	perl -i -0pe 's/\*struct\s\{\n\s\sAdditionalProperties\smap\[string\]string\s`json:"-"`\n\s}/map\[string\]string/gms' apis/stackgres/v1/sgcluster.gen.go

	go run sigs.k8s.io/controller-tools/cmd/controller-gen object paths=./apis/stackgres/v1/...
	rm apis/stackgres/v1/*_crd.yaml

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: lint
lint: fmt vet ## All-in-one linting
	@echo 'Check for uncommitted changes ...'
	git diff --exit-code

##@ Build

.PHONY: build
build: export CGO_ENABLED = 0
build: generate fmt vet  ## Build manager binary.
build:
	@echo "GOOS=$$(go env GOOS) GOARCH=$$(go env GOARCH)"
	go build -o $(BIN_FILENAME)

.PHONY: test
test: ## Run tests
	go test ./...

.PHONY: docker-build
docker-build:
	env CGO_ENABLED=0 GOOS=$(DOCKER_IMAGE_GOOS) GOARCH=$(DOCKER_IMAGE_GOARCH) \
		go build -o ${BIN_FILENAME}
	docker build -t ${GHCR_IMG} .

.PHONY: docker-push
docker-push: docker-build ## Push docker image with the manager.
	docker push ${GHCR_IMG}

.PHONY: clean
clean:
	rm -rf bin/ appcat .work/ docs/node_modules $docs_out_dir .public .cache apiserver.local.config apis/generated default.sock
