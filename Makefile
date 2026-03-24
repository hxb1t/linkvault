# Script for run golang service and database migrations
SERVER_DIR := server
APP_NAME := linkvault
GO := go
DOCKER := docker
API_ENTRYPOINT := ./cmd/api
BINARY := tmp/server

dev:
	cd $(SERVER_DIR) && $(GO) run air
run:
	cd $(SERVER_DIR) && $(GO) run $(API_ENTRYPOINT)
build:
	cd $(SERVER_DIR) && $(GO) build -o $(BINARY) $(MAIN)
deps:
	cd $(SERVER_DIR) && $(GO) mod download
test:
	cd $(SERVER_DIR) && $(GO) test ./... -v
docker/up:
	$(DOCKER) compose up --build -d
docker/down:
	$(DOCKER) compose down
docker/reset:
	$(DOCKER) compose down -v
	$(DOCKER) compose up --build -d

help:
	@echo "Usage: make [target]"
	@echo .
	@echo dev              run with hot reload (requires air)
	@echo run              run without hot reload
	@echo build            compile app binary
	@echo test             run all tests
	@echo deps             download all depedencies
	@echo docker/up        start all docker containers
	@echo docker/down      stop all docker containers
	@echo docker/reset     rebuild and restart all docker containers


.PHONY: dev run build deps test docker/up docker/down docker/reset
.DEFAULT_GOAL := help
