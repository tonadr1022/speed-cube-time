MODULE := $(shell go list -m)
MIGRATE := docker-compose run --rm migrate -path /migrations \
		   -database postgres://postgres:password@db:5432/cube?sslmode=disable
.PHONY: build
build: ## build the API server binary
	CGO_ENABLED=0 go build -o server $(MODULE)/cmd/server
	

.PHONY: build-docker
build-docker: ## build API server as a docker image
	docker build cmd/server/Dockerfile -t server .


.PHONY: up
up: ## run docker compose up
	docker compose up --build -d

.PHONY: stop
stop: ## stop docker compose
	docker compose stop

.PHONY: down
down: ## remove docker compose containers
	docker compose down

.PHONY: migrate
migrate:
	$(MIGRATE) up


.PHONY: migrate-reset
migrate-reset: ## reset database and re-run all migrations
	@echo "Resetting database..."
	$(MIGRATE) drop -f
	@echo "Running all database migrations..."
	@$(MIGRATE) up

.PHONY: test
test: ## run all unit tests
	@echo "mode: count" > coverage-all.out
	@$(foreach pkg,$(PACKAGES), \
		go test  -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)

.PHONY: test-cover
test-cover: test ## run unit tests and show test coverage information
	go tool cover -html=coverage-all.out

