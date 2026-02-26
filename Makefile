IMAGE ?= rafaelpm/api-stress-test
TAG ?= latest

.PHONY: help
help: ## Show this menu
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)


.PHONY: build
build: ## Build the container image
	@docker build -t $(IMAGE):$(TAG) .

.PHONY: run
run: ## Run the code locally
	@go run cmd/main.go --url=https://google.com --requests=10 --concurrency=2


.PHONY: runc
runc: ## Run the app via docker
	@docker run $(IMAGE):$(TAG) --url=https://google.com --requests=50 --concurrency=10

.PHONY: push
push: build ## Build and push container image
	@docker push $(IMAGE):$(TAG)

.PHONY: verify
verify: ## Validate tests and local/docker execution
	@go test ./...
	@$(MAKE) run
	@$(MAKE) build
	@$(MAKE) runc
