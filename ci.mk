# Image URL to use all building/pushing image targets
IMG_TAG ?= latest
APP_NAME ?= appcat
ORG ?= vshn
GHCR_IMG ?= ghcr.io/$(ORG)/$(APP_NAME):$(IMG_TAG)
DOCKER_CMD ?= docker

# For alpine image it is required the following env before building the application
DOCKER_IMAGE_GOOS = linux
DOCKER_IMAGE_GOARCH = amd64

.PHONY: docker-build
docker-build:
	env CGO_ENABLED=0 GOOS=$(DOCKER_IMAGE_GOOS) GOARCH=$(DOCKER_IMAGE_GOARCH) \
		go build -o ${BIN_FILENAME}
	docker build --platform $(DOCKER_IMAGE_GOOS)/$(DOCKER_IMAGE_GOARCH) -t ${GHCR_IMG} .

.PHONY: docker-build-branchtag
IMG_TAG =  $(shell git rev-parse --abbrev-ref HEAD | sed 's/\//_/g')
docker-build-branchtag: docker-build ## Build docker image with current branch name

.PHONY: docker-push
docker-push: docker-build ## Push docker image with the manager.
	docker push ${GHCR_IMG}

.PHONY: docker-push-branchtag
IMG_TAG =  $(shell git rev-parse --abbrev-ref HEAD | sed 's/\//_/g')
docker-push-branchtag: docker-build-branchtag docker-push ## Push docker image with current branch name

.PHONY: function-build
function-build: docker-build
	rm -f package/*.xpkg
	go run github.com/crossplane/crossplane/cmd/crank@v1.16.0 xpkg build -f package --verbose --embed-runtime-image=${GHCR_IMG} -o package/package.xpkg

.PHONY: function-push-package
function-push-package: function-build
	go run github.com/crossplane/crossplane/cmd/crank@v1.16.0 xpkg push -f package/package.xpkg ${GHCR_IMG}-func --verbose

.PHONY: function-build-branchtag
IMG_TAG =  $(shell git rev-parse --abbrev-ref HEAD | sed 's/\//_/g')
function-build-branchtag: docker-build-branchtag
	rm -f package/*.xpkg
	go run github.com/crossplane/crossplane/cmd/crank@v1.16.0 xpkg build -f package --verbose --embed-runtime-image=${GHCR_IMG} -o package/package.xpkg

.PHONY: function-push-package-branchtag
IMG_TAG =  $(shell git rev-parse --abbrev-ref HEAD | sed 's/\//_/g')
function-push-package-branchtag: function-build-branchtag
	go run github.com/crossplane/crossplane/cmd/crank@v1.16.0 xpkg push -f package/package.xpkg ${GHCR_IMG}-func --verbose
