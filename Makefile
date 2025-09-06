APP_NAME := cmit
GO := go
GOLANGCI := golangci-lint
BIN_DIR := bin
OS_LIST := linux darwin windows
ARCH := amd64

VERSION := $(shell git describe --tags --always --dirty)

ifeq ($(OS),)
	OS := $(shell uname | tr '[:upper:]' '[:lower:]')
endif

.PHONY: all
all: build

.PHONY: build
build:
ifeq ($(OS),Windows_NT)
	set GOOS=windows&& set GOARCH=amd64&& go build -ldflags="-X main.Version=$(VERSION)" -o $(BIN_DIR)/$(APP_NAME).exe ./cmd/cmit
else
	GOOS=$(OS) GOARCH=$(ARCH) go build -ldflags="-X main.Version=$(VERSION)" -o $(BIN_DIR)/$(APP_NAME) ./cmd/cmit
endif


.PHONY: build-all
build-all:
	@mkdir -p $(BIN_DIR)
	@for os in $(OS_LIST); do \
		ext=""; \
		if [ $$os = "windows" ]; then ext=".exe"; fi; \
		echo "Building for $$os..."; \
		GOOS=$$os GOARCH=$(ARCH) $(GO) build -ldflags="-X main.Version=$(VERSION)" -o $(BIN_DIR)/$(APP_NAME)_$$os$$ext ./cmd/cmit; \
	done

.PHONY: test
test:
	$(GO) test -v ./...

.PHONY: lint
lint:
	$(GOLANGCI) run ./...

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

.PHONY: ci
ci: lint test build

.PHONY: release
release: build
	@echo "Release built for $(OS)/$(ARCH) with version $(VERSION)"
	@mkdir -p dist
	@cp $(BIN_DIR)/$(APP_NAME)* dist/
