DOCKER_COMPOSE_FILE = docker-compose.yml
MYSQL_SERVICE = product-hub-mysql
GO_IMAGE = golang:1.23.1-alpine
CONTAINER_NAME = product-hub-app
APP_PORT = 10000

# Default target: Build and run the application
.PHONY: all
all: up run

# Start MySQL and other services
.PHONY: up
up:
	@docker compose -f $(DOCKER_COMPOSE_FILE) up -d

# Run the Go application in Docker
.PHONY: run
run:
	@docker run --rm -it -v $(shell pwd):/go/src/app -w /go/src/app -p $(APP_PORT):$(APP_PORT) $(GO_IMAGE) go run main.go

# Stop all services
.PHONY: stop
down:
	@docker compose -f $(DOCKER_COMPOSE_FILE) down

# Clean up Docker containers and images
.PHONY: clean
clean:
	@docker compose -f $(DOCKER_COMPOSE_FILE) down --volumes --rmi all
	@docker system prune -f

# Run Go tests
.PHONY: test
test:
	@docker run --rm -v $(shell pwd):/go/src/app -w /go/src/app $(GO_IMAGE) go test ./...
