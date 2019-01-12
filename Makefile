# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=_cabal-a

all: build
build: 
	$(GOBUILD) -o ./bin/$(BINARY_NAME) ./src
# test: 
# 	$(GOTEST) -v ./...
clean: 
	rm -f ./bin/$(BINARY_NAME)
run:
	$(GOBUILD) -o ./bin/$(BINARY_NAME)  ./src
	./bin/$(BINARY_NAME)
# deps:
# 	$(GOGET) github.com/markbates/goth
# 	$(GOGET) github.com/markbates/pop