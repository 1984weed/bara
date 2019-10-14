# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=bara
SERVER_FILE=server/server.go
BINARY_UNIX=$(BINARY_NAME)_unix

all: build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v $(SERVER_FILE) 

generate: 
	go run github.com/99designs/gqlgen -v
