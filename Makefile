.PHONY: all start stop build run test docker-build docker-push clean

all: build

start: ## Start the backend service and its dependencies using Docker Compose
	docker-compose up -d

stop: ## Stop the backend service and its dependencies using Docker Compose
	docker-compose down

build: ## Build the Go backend application
	go build -o any-api-backend ./cmd/api

run: build ## Run the Go backend application directly
	./any-api-backend

test: ## Run all Go tests
	go test ./...

docker-build: ## Build the Docker image for the backend
	docker build -t any-api-backend:latest .

docker-push: ## Build and push the Docker image to ECR using the build-and-push script. Usage: make docker-push TAG=v1.0.2
ifndef TAG
	$(error TAG is not set. Usage: make docker-push TAG=v1.0.2)
endif
	bash build-and-push.sh $(TAG)

clean: ## Clean up build artifacts
	rm -f any-api-backend
