# ref: https://github.com/dtan4/s3url/blob/master/Makefile
NAME := grapi-sample
VERSION := 1.0.0
REVISION := $(shell git rev-parse --short HEAD)

SRCS     := $(shell find . -type f -name '*.go')
LDFLAGS  := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""
NOVENDOR := $(shell go list ./... | grep -v vendor)

DIST_DIRS := find * -type d -exec

.DEFAULT_GOAL := help

.PHONY: help
help: ## Show help see: https://postd.cc/auto-documented-makefile/
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: dep
dep: ## Install dep
ifeq ($(shell command -v dep 2> /dev/null),)
	go get -u github.com/golang/dep/cmd/dep
endif

.PHONY: deps
deps: dep ## Install dependency libraries
	dep ensure -v

.PHONY: update-deps
update-deps: dep ## Update dependency libraries
	dep ensure -update -v

.PHONY: deps-vendor
deps-vendor: dep ## Install dependency libraries with only Gopkg.lock file. for CI
	dep ensure -v -vendor-only=true

.PHONY: dist
dist: ## Build app and archive binary
	cd dist && \
	$(DIST_DIRS) cp ../LICENSE {} \; && \
	$(DIST_DIRS) cp ../README.md {} \; && \
	$(DIST_DIRS) tar -zcf $(NAME)-$(VERSION)-{}.tar.gz {} \; && \
	$(DIST_DIRS) zip -r $(NAME)-$(VERSION)-{}.zip {} \; && \
	cd ..

.PHONY: build-linux
build-linux: ## Build app for linux arch
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o dist/$(NAME)

.PHONY: run
run: deps ## Run development server
ifeq ($(shell command -v gin 2> /dev/null),)
	go get -u github.com/codegangsta/gin
endif
	gin -a 3000 -p 3001 -i --notifications grapi server

.PHONY: protoc
protoc: ## Generate protc files using grapi
	grapi protoc
	find api -name *.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'
