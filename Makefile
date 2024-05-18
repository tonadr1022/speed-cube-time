MODULE := $(shell go list -m)

.PHONY: build
build: ## build the API server binary
	CGO_ENABLED=0 GOOS=linux go build -o /speed-cube-time $(MODULE)/cmd/server
	

.PHONY: build-docker
build-docker:
	docker build -f cmd/server/Dockerfile

