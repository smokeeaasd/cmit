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
	set GOOS=windows&& set GOARCH=$(ARCH)&& $(GO) build -ldflags="-X 'github.com/smokeeaasd/cmit/internal/version.Version=$(VERSION)'" -o $(BIN_DIR)/$(APP_NAME)-windows-$(ARCH).exe ./cmd/main
else
	GOOS=$(OS) GOARCH=$(ARCH) $(GO) build -ldflags="-X 'github.com/smokeeaasd/cmit/internal/version.Version=$(VERSION)'" -o $(BIN_DIR)/$(APP_NAME)-$(OS)-$(ARCH) ./cmd/main
endif

.PHONY: build-all
build-all:
	@mkdir -p $(BIN_DIR)
	@for os in $(OS_LIST); do \
		ext=""; \
		if [ $$os = "windows" ]; then ext=".exe"; fi; \
		echo "Building for $$os..."; \
		GOOS=$$os GOARCH=$(ARCH) $(GO) build -ldflags="-X 'github.com/smokeeaasd/cmit/internal/version.Version=$(VERSION)'" -o $(BIN_DIR)/$(APP_NAME)-$$os-$(ARCH)$$ext ./cmd/main; \
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
