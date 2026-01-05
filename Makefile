.PHONY: help run build test test-coverage clean migrate lint swagger deps tidy dev

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run: ## Run the application
	@go run main.go

build: ## Build the application
	@echo "Building..."
	@go build -o bin/api main.go

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

test-fail: ## Run tests and show failed test
	@echo "Running tests and showing failed tests (failures only)..."
	@tmp=$$(mktemp); \
	if command -v ginkgo >/dev/null 2>&1 && command -v jq >/dev/null 2>&1; then \
		ginkgo -r --no-color --keep-going --fail-fast --silence-skips --json-report $$tmp >/dev/null; \
		status=$$?; \
		jq -r '.suite.suiteDescription as $$suite | .specReports[] | select(.state=="failed") | "" + $$suite + "\n" + .leafNodeText + "\n" + .failure.message + "\nLocation: " + .leafNodeLocation.fileName + ":" + (.leafNodeLocation.lineNumber|tostring) + "\n---"' $$tmp; \
		rm -f $$tmp; \
		[ $$status -eq 0 ] && echo "All specs passed."; \
		exit $$status; \
	elif command -v ginkgo >/dev/null 2>&1; then \
		ginkgo -r --no-color --keep-going --fail-fast --silence-skips --succinct; \
	else \
		go test ./...; \
	fi

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -v -race -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html



deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod verify

tidy: ## Tidy dependencies
	@echo "Tidying dependencies..."
	@go mod tidy

dev: ## Run with hot reload (requires air)
	@echo "Starting development server with hot reload..."
	@air

swagger: ## Generate swagger documentation
	@echo "Generating swagger documentation..."
	@swag init
	@echo "✓ Swagger documentation generated in docs/"
