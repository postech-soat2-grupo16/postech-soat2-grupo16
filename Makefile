.PHONY: help
db-run:
	@docker-compose up -d fastfood_db
app-run:
	@docker-compose up fastfood_app
build-docker:
	@docker-compose build
run-all:
	@docker-compose up