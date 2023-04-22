VERSION     := $(shell git describe --tags --always --abbrev=8)
SYSTEM      :=
IMPORT_PATH := github.com/milgradesec/sensord
CGO_ENABLED := 0
DATE        := $(shell date -u '+%Y-%m-%d-%H%M UTC')
BUILDFLAGS  := -v -ldflags="-s -w -X main.Version=$(VERSION)"

.PHONY: all
all: build

.PHONY: clean
clean:
	go clean
	rm -f sensord

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
	./sensord -dev

.PHONY: install
install:
	cp sensord /usr/bin/sensord