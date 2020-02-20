# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=bara
SANDBOX_BINARY_NAME=sandbox-cli
SERVER_FILE=server/server.go
SANDBOX_MAIN_FILE=sandbox/main.go
BINARY_UNIX=$(BINARY_NAME)_unix

build: install build-server build-sandbox
install: 
	go get ./...

build-server: 
	$(GOBUILD) -o $(BINARY_NAME) -v $(SERVER_FILE) 

build-sandbox: 
	$(GOBUILD) -o $(SANDBOX_BINARY_NAME) -v $(SANDBOX_MAIN_FILE) 

generate: 
	@go run github.com/99designs/gqlgen -v

integration-test: 
	@go test -v ./...

unit-test: 
	@go test -v -short ./...

docker-up-for-test:
	@docker-compose up -d bara.db-test

docker-local:
	@docker-compose up -d bara.db bara.redis

init-db-data:
	./scripts/init-data.sh
