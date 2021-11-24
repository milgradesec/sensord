VERSION     := $(shell git describe --tags --always --abbrev=0)
SYSTEM      :=
BUILDFLAGS  := -v
CGO_ENABLED := 0

.PHONY: all
all: build

.PHONY: clean
clean:
	go clean

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build:
	CGO_ENABLED=$(CGO_ENABLED) $(SYSTEM) go build $(BUILDFLAGS)

.PHONY: run
run:
	sudo ./sensord
