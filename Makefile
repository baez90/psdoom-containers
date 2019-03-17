VERSION = $(shell git describe --dirty --tags --always)
REPO = github.com/baez90/psdoom-containers
BUILD_PATH = $(REPO)/cmd/psdoom-containers
PKGS = $(shell go list ./... | grep -v /vendor/)
TEST_PKGS = $(shell find . -type f -name "*_test.go" -printf '%h\n' | sort -u)
GOARGS = GOOS=linux GOARCH=amd64
BINARY_NAME = psdoom-containers
DIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
DEBUG_PORT = 2345

export CGO_ENABLED:=0
export WATCH_NAMESPACE=default

.PHONY: all clean clean-all clean-vendor rebuild format revive generate dep dep-check dep-check test integration-test compile run-local debug-local watch watch-test docs serve-docs ensure-dep ensure-revive ensure-reflex ensure-delve

all: format compile

clean-all: clean clean-vendor

rebuild: clean format compile

install:
	@cp -f $(DIR)/$(BINARY_NAME) $(GOPATH)/bin/

format:
	@go fmt $(PKGS)

generate:
	@mkdir -p $(DIR)/internal/pkg/api/k8s/generated
	@protoc -I ./api/ ./api/k8s-daemon.proto --go_out=plugins:./internal/pkg/api/k8s/generated

revive: ensure-revive
	@revive --config $(DIR)assets/lint/config.toml -exclude $(DIR)vendor/... -formatter friendly $(DIR)...

dep: ensure-dep
	@dep ensure -v

dep-update: ensure-dep
	@dep ensure -update -v

dep-check: ensure-dep
	@dep check

dep-graph:
	@dep status -dot | dot -T png > /tmp/dep-graph.png && feh /tmp/dep-graph.png&

clean:
	rm -f debug $(BINARY_NAME)

clean-vendor:
	rm -rf vendor/

test:
	@go test -v $(TEST_PKGS)

integration-test:
	@go test -v $(TEST_PKGS) -tags=integration

compile:
	@$(GOARGS) go build -o $(BINARY_NAME) $(BUILD_PATH)

watch-test: ensure-reflex
	@reflex -r '_test\.go$$' -s -- sh -c 'make test'

serve-docs: ensure-reflex docs
	@reflex -r '\.md$$' -s -- sh -c 'mdbook serve -d $(DIR)/public -n 127.0.0.1 $(DIR)/docs'

docs:
	@mdbook build -d $(DIR)/public $(DIR)/docs

ensure-dep:
ifeq (, $(shell which dep))
	$(shell go get -u github.com/golang/dep/cmd/dep)
endif

ensure-revive:
ifeq (, $(shell which revive))
	$(shell go get -u github.com/mgechev/revive)
endif

ensure-delve:
ifeq (, $(shell which dlv))
	$(shell go get -u github.com/go-delve/delve/cmd/dlv)
endif

ensure-reflex:
ifeq (, $(shell which reflex))
	$(shell go get -u github.com/cespare/reflex)
endif
