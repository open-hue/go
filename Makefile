# Go
GO ?= `which go`

# Git
GIT ?= `which git`

# The specification file from the open-hue/spec sub-module
SPEC_FILE ?= spec/spec.yaml

.PHONY: help
help: ## Lists help commands
	@grep -h -E '^[a-zA-Z_\-\/]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-36s\033[0m %s\n", $$1, $$2}'

.PHONY: spec/update
spec/update: ## Refreshes the OpenAPI specification from its source
	@$(GIT) submodule update --init --remote spec

.PHONY: build
build: generate ## Builds client code
	@$(GO) build

.PHONY: generate
generate: ## Generates client code from OpenAPI specification
	@$(GO) generate

.PHONY: deps/update
deps/update: ## Updates go dependencies
	@$(GO) get -u
	@$(GO) mod tidy
