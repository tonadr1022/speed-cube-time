MODULE := $(shell go list -m)

CONFIG_FILE ?= ./config/local.yaml
APP_DSN ?= $(shell sed -n 's/^dsn:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG_FILE))
MIGRATE_OLD := docker run -v $(shell pwd)/migrations:/migrations \
		   --network host migrate/migrate:v4.17.1 \
		   -path=/migrations/ \
		   -database "$(APP_DSN)"
MIGRATE := docker-compose run --rm migrate -path /migrations \
		-database "$(APP_DSN)" 

.PHONY: build
build: ## build the API server binary
	CGO_ENABLED=0 GOOS=linux go build -a -o server $(MODULE)/cmd/server
	

.PHONY: build-docker
build-docker: ## build API server as a docker image
	docker build -f cmd/server/Dockerfile -t server .

.PHONY: run 
run: 
	go run cmd/server/main.go

.PHONY: up
up: ## run docker compose up
	docker compose up --build -d

.PHONY: stop
stop: ## stop docker compose
	docker compose stop

.PHONY: down
down: ## remove docker compose containers
	docker compose down

.PHONY: db-start
db-start: ## start the database server
	docker run -d --rm \
		--name roach \
		--env COCKROACH_DB=speed_cube_time \
		--env COCKROACH_USER=username \
		--env COCKROACH_PASSWORD=roach \
		--env PGPORT=26257 \
		-p 26257:26257\
		-p 8080:8080 \
		-v roach:/cockroach/cockroach-data \
		cockroachdb/cockroach:latest-v23.2 start-single-node \
		--insecure

.PHONY: db-stop
db-stop: ## stop the database server
	docker stop roach


.PHONY: migrate
migrate:
	$(MIGRATE) up

.PHONY: migrate-new
migrate-new: ## create a new database migration
	@read -p "Enter the name of the new migration: " name; \
	$(MIGRATE) create -ext sql -dir /migrations/ $${name// /_}

.PHONY: migrate-reset
migrate-reset: ## reset database and re-run all migrations
	@echo "Resetting database..."
	$(MIGRATE) drop -f
	@echo "Running all database migrations..."
	@$(MIGRATE) up
