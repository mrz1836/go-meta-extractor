# Common makefile commands & variables between projects
include .make/common.mk

# Common Golang makefile commands & variables between projects
include .make/go.mk

## Not defined? Use default repo name which is the application
ifeq ($(REPO_NAME),)
	REPO_NAME="go-meta-extractor"
endif

## Not defined? Use default repo owner
ifeq ($(REPO_OWNER),)
	REPO_OWNER="mrz1836"
endif

.PHONY: clean install-tools

all: ## Runs multiple commands
	@$(MAKE) test

clean: ## Remove previous builds and any cached data
	@echo "cleaning local cache..."
	@go clean -cache -testcache -i -r
	@test $(DISTRIBUTIONS_DIR)
	@if [ -d $(DISTRIBUTIONS_DIR) ]; then rm -r $(DISTRIBUTIONS_DIR); fi

install-tools: ## Install the go tools
	@echo "installing local Go tools..."
	@cd tools && go install $(shell cd tools && go list -f '{{ join .Imports " " }}' -tags=tools)

install-all-contributors: ## Installs all contributors locally
	@echo "installing all-contributors cli tool..."
	@yarn global add all-contributors-cli

release:: ## Runs common.release then runs godocs
	@$(MAKE) godocs

update-contributors: ## Regenerates the contributors html/list
	@echo "generating contributor html..."
	@all-contributors generate

update-tools: ## Update all go tools
	@echo "updating local tool dependencies..."
	@cd tools && go get -u -v && go mod tidy -v
