# GoTemplate CLI — Build System
# Usage: make [build|test|run|clean|lint|all]

BINARY   := gotemplate-cli
VERSION  := 0.1.0
LDFLAGS  := -ldflags "-X main.version=$(VERSION)"
GO       := go
GOFLAGS  := -v

.PHONY: all build test run clean lint help

## all: Build the binary and run tests
all: build test

## build: Compile the binary
build:
	$(GO) build $(GOFLAGS) $(LDFLAGS) -o $(BINARY) .

## test: Run all tests
test:
	$(GO) test $(GOFLAGS) ./...

## run: Build and launch the CLI help
run: build
	./$(BINARY) --help

## testv: Run all tests with verbose output and coverage
testv:
	$(GO) test -v -cover ./...

## clean: Remove build artifacts
clean:
	rm -f $(BINARY)
	$(GO) clean -cache

## lint: Run go vet and gofmt linting
lint:
	$(GO) vet ./...
	gofmt -l -s .

## deps: Download and tidy module dependencies
deps:
	$(GO) mod download
	$(GO) mod tidy

## install: Build and install to GOPATH/bin
install:
	$(GO) install $(LDFLAGS) .

## fmt: Format all Go source files
fmt:
	$(GO) fmt ./...

## help: Display this help
help:
	@echo "GoTemplate CLI — Available targets:"
	@echo "  build    Compile the binary"
	@echo "  test     Run all tests"
	@echo "  run      Build and run --help"
	@echo "  testv    Run tests verbose with coverage"
	@echo "  clean    Remove build artifacts"
	@echo "  lint     Run go vet and gofmt"
	@echo "  deps     Download and tidy dependencies"
	@echo "  install  Build and install to GOPATH/bin"
	@echo "  fmt      Format all Go files"
	@echo "  all      Build and test"
