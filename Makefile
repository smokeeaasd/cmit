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

.PHONY: all build build-linux build-macos build-windows build-all test lint clean ci release

all: build

# Build para o SO atual
build:
ifeq ($(OS),Windows_NT)
	set GOOS=windows&& set GOARCH=$(ARCH)&& $(GO) build -ldflags="-X 'github.com/smokeeaasd/cmit/internal/version.Version=$(VERSION)'" -o $(BIN_DIR)/$(APP_NAME)-windows-$(ARCH).exe ./cmd/main
else
	GOOS=$(OS) GOARCH=$(ARCH) $(GO) build -ldflags="-X 'github.com/smokeeaasd/cmit/internal/version.Version=$(VERSION)'" -o $(BIN_DIR)/$(APP_NAME)-$(OS)-$(ARCH) ./cmd/main
endif

# Builds individuais por SO
build-linux:
	GOOS=linux GOARCH=$(ARCH) $(GO) build -ldflags="-X 'github.com/smokeeaasd/cmit/internal/version.Version=$(VERSION)'" -o $(BIN_DIR)/$(APP_NAME)-linux-$(ARCH) ./cmd/main

build-darwin:
	GOOS=darwin GOARCH=$(ARCH) $(GO) build -ldflags="-X 'github.com/smokeeaasd/cmit/internal/version.Version=$(VERSION)'" -o $(BIN_DIR)/$(APP_NAME)-darwin-$(ARCH) ./cmd/main

build-windows:
	GOOS=windows GOARCH=$(ARCH) $(GO) build -ldflags="-X 'github.com/smokeeaasd/cmit/internal/version.Version=$(VERSION)'" -o $(BIN_DIR)/$(APP_NAME)-windows-$(ARCH).exe ./cmd/main

# Build para todos
build-all: build-linux build-darwin build-windows

# Testes
test:
	$(GO) test -v ./...

# Lint
lint:
	$(GOLANGCI) run ./...

# Limpeza
clean:
	rm -rf $(BIN_DIR) dist

# CI
ci: lint test build

# Release: copia binários para dist
release: build-all
	@echo "Release built with version $(VERSION)"
	@mkdir -p dist
	@cp $(BIN_DIR)/$(APP_NAME)* dist/
