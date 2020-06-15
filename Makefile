# Go parameters
.PHONY:  testall test testl testv coverage threshold lint run depgraph protos
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GORUN=$(GOCMD) run .
GOCOV=$(GOCMD) tool cover -html=coverage.out
GOTEST=$(GOCMD) test -tags test
GOGET=$(GOCMD) get
GODEP=godepgraph -s -o  github.com/webmalc/namecheap-dns github.com/webmalc/namecheap-dns | dot -Tpng -o godepgraph.png
BINARY_NAME=ndns.app

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	GOENV=test $(GOTEST) ./... -coverprofile=coverage.out

testv:
	GOENV=test $(GOTEST) -v ./... -coverprofile=coverage.out

depgraph:
	$(GODEP)

coverage:
	$(GOCOV)

threshold:
	overcover --coverprofile coverage.out --threshold 90 --summary
testl: testv lint

testall: test lint threshold

protos:
	protoc -I protos/ protos/changer.proto --go_out=plugins=grpc:protos/changer

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

lint:
	golangci-lint run ./...
	golint ./...
run:
	$(GORUN)
