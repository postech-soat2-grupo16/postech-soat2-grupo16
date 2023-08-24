.PHONY: help
help: ## Display this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: db-run
db-run: ## Run the database container
	@docker-compose up -d fastfood_db

.PHONY: app-run
app-run: ## Run the application container
	@docker-compose up -d fastfood_app

.PHONY: build-all
build-all: ## Build docker images
	@docker-compose build

.PHONY: run-all
run-all: ## Run all containers
	@docker-compose up

.PHONY: kill-all
kill-all: ## Run all containers
	@docker-compose down --volumes --remove-orphans

.PHONY: db-reset
db-reset: ## Reset table registers to initial state (with seeds)
	@docker exec -u postgres fastfood_db psql fastfood_db postgres -f /migration/seeds/seeds.sql

.PHONY: test
test: db-reset ## Execute the tests in the development environment
	@go test ./... -count=1 -race -timeout 2m

.PHONY: lint
lint: ## Execute syntatic analysis in the code and autofix minor problems
	@golangci-lint run --fix

.PHONY: ci
ci: lint test ## Execute the tests and lint commands

.PHONY: db-logs
db-logs: ## Show database logs
	@docker logs -f --tail 100 fastfood_db

.PHONY: app-logs
app-logs: ## Show application logs
	@docker logs -f --tail 100 fastfood_app

.PHONY: update-docs
update-docs: ## Update swagger docs
	@swag init -d ./  --parseDependency --parseInternal --parseDepth 2 -o docs

.PHONY: k8s-apply
k8s-apply: ## Apply kubernetes resources
	@kubectl apply -f k8s.yml

.PHONY: get-pod
get-pod:
	@export POD_POSTGRES=$$(kubectl get pods -l app=postgres -o jsonpath='{.items[0].metadata.name}')

.PHONY: exec-migration
exec-migration: get-pod ## Execute migration in kubernetes
	@kubectl exec -it $$POD_POSTGRES -- psql -U postgres -d postgres -f ./migration/init.sql
	@kubectl exec -it $$POD_POSTGRES -- psql -U postgres -d postgres -f ./migration/migration.sql
