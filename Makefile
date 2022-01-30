VERSION     := $(shell git describe --tags --always --abbrev=0)
SYSTEM      :=
BUILDFLAGS  := -v -ldflags="-s -w -X main.Version=$(VERSION)"
IMPORT_PATH := github.com/milgradesec/sensord
CGO_ENABLED := 0

.PHONY: all
all: build

.PHONY: clean
clean:
	go clean
	rm sensord

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build:
	CGO_ENABLED=$(CGO_ENABLED) $(SYSTEM) go build $(BUILDFLAGS) $(IMPORT_PATH)/cmd/sensord


.PHONY: run
run:
	./sensord

.PHONY: install
install:
	cp sensord /usr/bin/sensord
