BINS = server

VERSION ?= $(shell git describe --tags --always --dirty)

LD_FLAGS = -X $(shell go list -m)/pkg/version.Version=$(VERSION)

FLAGS = -race

DEBUG ?= 0

ifeq ($(DEBUG), 1)
	FLAGS += -gcflags "all=-N -l" -ldflags="$(LD_FLAGS)"
else
	FLAGS += -gcflags "all=-trimpath=$(pwd)" -asmflags "all=-trimpath=$(pwd)" -ldflags="-s -w $(LD_FLAGS)"
endif

all: # @HELP Build the project
all: build

test: # @HELP Run all tests
test: vet lint
	go test -race -coverprofile cover.out ./... |&pp

test-ci: # @HELP Run tests and output JSON format suitable for CI tool
test-ci: vet lint
	go test -race -coverprofile cover.out -json ./...

coverage: # @HELP Generate HTML coverage report
coverage:
	go tool cover -html=cover.out -o coverage.html

vet: # @HELP Run vet tool against code
vet:
	go vet ./...

lint: # @HELP runs golangci-lint
lint:
	golangci-lint run

outdir:
	@mkdir -p ./bins

build: outdir
	@for bin in $(BINS); do \
		go build $(FLAGS) -o ./bins/$$bin ./cmd/$$bin; \
	done

clean:
	@rm -rf bins/*

version: # @HELP outputs the version string
version:
	echo $(VERSION)

.PHONY: clean help test

help: # @HELP prints this message
help:
	@echo "VARIABLES:"
	@echo "  BINS = $(BINS)"
	@echo
	@echo "TARGETS:"
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST)     \
	    | awk '                                   \
	        BEGIN {FS = ": *# *@HELP"};           \
	        { printf "  %-30s %s\n", $$1, $$2 };  \
	    '
