
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_LOC=bin
CAT_BINARY_NAME=cat
UNIQ_BINARY_NAME=uniq
PROJECT_HOME=$(shell pwd)

all: test build
build: 
	$(GOBUILD) -o ./$(BINARY_LOC)/$(CAT_BINARY_NAME) -v ./cmd/$(CAT_BINARY_NAME)/...
	$(GOBUILD) -o ./$(BINARY_LOC)/$(UNIQ_BINARY_NAME) -v ./cmd/$(UNIQ_BINARY_NAME)/...
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -rf $(BINARY_LOC)