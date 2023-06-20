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

db-reset: ## Reset table registers to initial state (with seeds)
	@docker exec -u postgres fastfood_db psql fastfood_db postgres -f /migration/seeds/seeds.sql
.PHONY: test
test: db-reset ## Execute the tests in the development environment
	@go test ./... -count=1 -race -timeout 2m

.PHONY: db-logs
db-logs: ## Show database logs
	@docker logs -f --tail 100 fastfood_db

.PHONY: app-logs
app-logs: ## Show application logs
	@docker logs -f --tail 100 fastfood_app