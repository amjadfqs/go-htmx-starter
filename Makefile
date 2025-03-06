.PHONY: dev build clean run test generate css css-watch air-install air setup help

# Default target
.DEFAULT_GOAL := help

# Variables
APP_NAME := go-htmx-starter
MAIN_PATH := ./cmd/server/main.go
BUILD_DIR := ./build
STATIC_DIR := ./static
CSS_INPUT := $(STATIC_DIR)/css/input.css
CSS_OUTPUT := $(STATIC_DIR)/css/style.css

# Development server with hot reload
dev: ## Run development server with Air (hot reload)
	air

# Build the application
build: generate css ## Build the application
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)

# Clean build artifacts
clean: ## Clean build artifacts
	rm -rf $(BUILD_DIR)
	rm -f $(CSS_OUTPUT)
	@echo "Cleaned build artifacts"

# Run the application
run: build ## Run the compiled application
	$(BUILD_DIR)/$(APP_NAME)

# Run tests
test: ## Run tests
	go test -v ./...

# Generate templ templates
generate: ## Generate templ templates
	templ generate

# Generate CSS once
css: ## Compile Tailwind CSS once
	npx @tailwindcss/cli -i $(CSS_INPUT) -o $(CSS_OUTPUT)

# Watch CSS files for changes
css-watch: ## Watch and compile Tailwind CSS
	npx @tailwindcss/cli -i $(CSS_INPUT) -o $(CSS_OUTPUT) --watch

# Install Air for hot reloading
air-install: ## Install Air hot reload tool
	go install github.com/air-verse/air@latest
	@echo "Air installed successfully"

# Setup the project
setup: ## Setup the project dependencies
	go mod tidy
	go get github.com/a-h/templ/cmd/templ@latest
	go install github.com/a-h/templ/cmd/templ@latest
	npm install tailwindcss @tailwindcss/cli
	@echo "Project dependencies installed"

# Start development with concurrent processes
dev-all: ## Run all development processes concurrently (requires concurrently npm package)
	concurrently "make generate" "make css-watch" "make dev"

# Help command
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
