.PHONY: clean build

GORUN = go run
GOBUILD = go build

all: build

clean:
	rm -rf bin

build: clean
	$(GOBUILD) -o bin/mclogger main.go

