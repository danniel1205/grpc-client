# Image URL to use all building/pushing image targets
IMG ?= danielguo/grpc-client:latest

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Run tests
test: fmt
	go test ./... -coverprofile cover.out

# Run go fmt against code
fmt:
	gofmt -w .
	goimports -w .

# Build the docker image
docker-build: test
	docker build . -t ${IMG}