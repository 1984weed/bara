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

all: build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v $(SERVER_FILE) 

build-sandbox: 
	$(GOBUILD) -o $(SANDBOX_BINARY_NAME) -v $(SANDBOX_MAIN_FILE) 

generate: 
	go run github.com/99designs/gqlgen -v