# This file is the development Makefile for the project in this repository.
# All variables listed here are used as substitution in these Makefile targets.
SERVICE_NAME = grail-participant-registry

.PHONY: help
help: ## Displays the Makefile help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

################################################################################

.PHONY: setup
setup: ## Downloads and install various libs for development.
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/vektra/mockery/v2/.../

.PHONY: build
build: lint ## Builds project binary.
	go build -ldflags "-w -s -X grail-participant-registry/internal/app.version=`git describe --tags --dirty` -X grail-participant-registry/internal/app.commitHash=`git rev-parse HEAD`" -race -o ./bin/$(SERVICE_NAME) -v ./cmd/app/app.go

.PHONY: test
test: lint ## Runs the test suite. Some projects might rely on a local development infrasructure to run tests. See `infra-up`.
	go test -v -race -bench=./... -benchmem -timeout=120s -cover -coverprofile=./test/coverage.txt ./...

.PHONY: test-ci
test-ci: lint ## Runs the test suite without -race because it is not supported on alpine
	# This should also run -race but it doesn't work on alpine
	go test -v -benchmem -timeout=120s -cover -coverprofile=./test/coverage.txt -bench=./... ./...

.PHONY: test-quick ## Run the quick test suite
test-quick:
	go test -short -failfast

.PHONY: run
run: build ## Builds project binary and executes it.
	bin/$(SERVICE_NAME)

.PHONY: full
full: clean build fmt lint test ## Cleans up, builds the service, reformats code, lints and runs the test suite.


################################################################################

.PHONY: docker-up
docker-up: docker/dev/.env ## Sets up local development infrastructure through Docker Compose.
	docker-compose -p $(SERVICE_NAME) --project-directory=docker/dev -f docker/dev/docker-compose.yml up -d

.PHONY: docker-down
docker-down: ## Tears down the local development infrastructure.
	docker-compose -p $(SERVICE_NAME) --project-directory=docker/dev -f docker/dev/docker-compose.yml down

################################################################################

init-env: ## Copy the default env file into place if no env file exists
	cp ./.env.dist ./.env
	cp ./docker/dev/.env.dist ./docker/dev/.env

################################################################################

.PHONY: lint
lint: ## Runs linter against the service codebase.
	golangci-lint run --config configs/golangci.yml
	@echo "[✔] \e[1m\e[32mLinter passed\e[0m"


.PHONY: fmt
fmt: ## Runs gofmt against the service codebase.
	gofmt -w -s .
	goimports -w .
	go clean ./...

.PHONY: tidy
tidy: ## Runs go mod tidy against the service codebase.
	go mod tidy

.PHONY: clean
clean: ## Removes temporary files and deletes the service binary.
	go clean ./...
	rm -f bin/$(SERVICE_NAME)

.PHONY: env
env: ## Displays the current Go environment variables.
	go env

.PHONY: doc
doc: ## Launches the godoc server locally. You can access it through a browser at the URL `http://localhost:8080/`.
	godoc -http=:8080 -index

.PHONY: version
version: ## Displays the current version of the Go toolchain.
	go version
