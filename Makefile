COVER_PKG=$$(go list ./... )$
# cover packages to use as go -coverpkg flag argument, separated by coma
COVER_PKG_FLAG=$(shell echo ${COVER_PKG} | sed 's/ /,/g')
REPOSITORY ?=
SERVICE_NAMESPACE = "go/std-server"
BRANCH ?= $$(git rev-parse --abbrev-ref HEAD)
TAG ?= $$(echo $(BRANCH) | tr '/' '-')
IMAGE = ${REPOSITORY}/${SERVICE_NAMESPACE}:${TAG}
PLATFORM_SUPPORT_ERROR := $(shell docker build -h 2>/dev/null | grep -q platform; echo $$?)
PLATFORM_FLAG :=
ifeq ($(PLATFORM_SUPPORT_ERROR),0)
	PLATFORM_FLAG += --platform=linux/amd64
endif

run:
	@go run cmd/main.go
# build image
build:
	@echo "Building ${SERVICE_NAMESPACE}..."
	@docker image build -t ${IMAGE} . --build-arg http_proxy --build-arg https_proxy ${PLATFORM_FLAG}
