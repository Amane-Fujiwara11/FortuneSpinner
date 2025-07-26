# Makefile for FortuneSpinner

# Variables
DOCKER_COMPOSE = docker-compose
DOCKER_COMPOSE_DEV = docker-compose -f docker-compose.dev.yml

# Colors for output
GREEN = \033[0;32m
YELLOW = \033[0;33m
RED = \033[0;31m
NC = \033[0m # No Color

.PHONY: help
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  ${GREEN}%-15s${NC} %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: dev
dev: ## Start development environment (fast, with hot reload)
	@echo "${GREEN}Starting development environment...${NC}"
	$(DOCKER_COMPOSE_DEV) up -d
	@echo "${GREEN}Development environment started!${NC}"
	@echo "Frontend: http://localhost:3000"
	@echo "Backend: http://localhost:8080"

.PHONY: dev-logs
dev-logs: ## Show development environment logs
	$(DOCKER_COMPOSE_DEV) logs -f

.PHONY: prod
prod: ## Start production environment
	@echo "${GREEN}Starting production environment...${NC}"
	$(DOCKER_COMPOSE) up -d --build
	@echo "${GREEN}Production environment started!${NC}"

.PHONY: stop
stop: ## Stop all containers
	@echo "${YELLOW}Stopping containers...${NC}"
	$(DOCKER_COMPOSE_DEV) stop
	$(DOCKER_COMPOSE) stop

.PHONY: down
down: ## Stop and remove all containers
	@echo "${YELLOW}Removing containers...${NC}"
	$(DOCKER_COMPOSE_DEV) down
	$(DOCKER_COMPOSE) down

.PHONY: clean
clean: down ## Clean everything (containers, volumes, images)
	@echo "${RED}Cleaning everything...${NC}"
	$(DOCKER_COMPOSE_DEV) down -v --rmi all
	$(DOCKER_COMPOSE) down -v --rmi all

.PHONY: migrate
migrate: ## Run database migrations
	@echo "${GREEN}Running migrations...${NC}"
	cd migrations && ./migrate.sh

.PHONY: backend-dev
backend-dev: ## Run backend locally (requires MySQL running)
	@echo "${GREEN}Starting backend locally...${NC}"
	cd backend && go run main.go

.PHONY: frontend-dev
frontend-dev: ## Run frontend locally
	@echo "${GREEN}Starting frontend locally...${NC}"
	cd frontend && npm start

.PHONY: build
build: ## Build all Docker images
	@echo "${GREEN}Building Docker images...${NC}"
	$(DOCKER_COMPOSE) build --parallel

.PHONY: rebuild
rebuild: ## Force rebuild all Docker images
	@echo "${GREEN}Rebuilding Docker images...${NC}"
	$(DOCKER_COMPOSE) build --no-cache --parallel

.PHONY: mysql
mysql: ## Connect to MySQL container
	docker exec -it fortunespinner-mysql mysql -uroot -prootpassword fortunespinner

.PHONY: backend-logs
backend-logs: ## Show backend logs
	docker logs -f fortunespinner-backend

.PHONY: frontend-logs
frontend-logs: ## Show frontend logs
	docker logs -f fortunespinner-frontend

.PHONY: mysql-logs
mysql-logs: ## Show MySQL logs
	docker logs -f fortunespinner-mysql

.PHONY: test-backend
test-backend: ## Run backend tests
	cd backend && go test ./...

.PHONY: test-frontend
test-frontend: ## Run frontend tests
	cd frontend && npm test

.PHONY: lint-backend
lint-backend: ## Run backend linter
	cd backend && go fmt ./...

.PHONY: ps
ps: ## Show running containers
	@docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"