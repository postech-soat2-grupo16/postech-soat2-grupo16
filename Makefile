.PHONY: help
help: ## Display this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: db-run
db-run: ## Run the database container
	@docker-compose up -d fastfood_db

.PHONY: app-run
app-run: ## Run the application container
	@docker-compose up fastfood_app

.PHONY: build-docker
build-docker: ## Build docker images
	@docker-compose build

.PHONY: run-all
run-all: ## Run all containers
	@docker-compose up

.PHONY: db-reset
db-reset: ## Reset table registers to initial state (with seeds)
	@echo "TODO: Implement"

.PHONY: test
test: ## Execute the tests in the development environment
	@go test ./... -count=1 -race -timeout 2m
