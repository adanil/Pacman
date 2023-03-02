PROJECT_PATH=$(CURDIR)
GO?=$(shell which go)
APP?=pacman
GOBIN?=$(CURDIR)/bin

$(value $(shell [ ! -d "$(CURDIR)/bin" ] && mkdir -p "$(CURDIR)/bin"))

GOLANGCI_BIN:=$(GOBIN)/golangci-lint
GOLANGCI_REPO=https://github.com/golangci/golangci-lint
GOLANGCI_LATEST_VERSION:= $(shell git ls-remote --tags --refs --sort='v:refname' $(GOLANGCI_REPO)|tail -1|egrep -o "v[0-9]+.*")
ifneq ($(wildcard $(GOLANGCI_BIN)),)
	GOLANGCI_CUR_VERSION=v$(shell $(GOLANGCI_BIN) --version|sed -E 's/.* version (.*) built from .* on .*/\1/g')
else
	GOLANGCI_CUR_VERSION=
endif

.PHONY: .golangci-lint-install
.golangci-lint-install: ##install linter tool
ifeq ($(filter $(GOLANGCI_CUR_VERSION), $(GOLANGCI_LATEST_VERSION)),)
	$(info Installing GOLANGCI-LINT $(GOLANGCI_LATEST_VERSION)...)
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s $(GOLANGCI_LATEST_VERSION)
	@chmod +x $(GOLANGCI_BIN)
else
	@echo 1 >/dev/null
endif

.PHONY: lint-local
lint-local: .golangci-lint-install ##run code full lint
	@echo perform full lint... && \
	$(GOLANGCI_BIN) cache clean && \
	$(GOLANGCI_BIN) run --config=$(CURDIR)/.golangci.yaml -v $(CURDIR)/...

.PHONY: lint-local-fast
lint-local-fast: .golangci-lint-install ##run code lint fastly
	@$(GOLANGCI_BIN) run --fast

.PHONY: go-deps
go-deps: ##check module dependencies
	@echo checking \"go\" modules deps ... && \
	$(GO) mod tidy && \
 	$(GO) mod vendor && \
	$(GO) mod verify && \
	echo -=OK=-

.PHONY: build
build: ##build the app
	@$(MAKE) go-deps && \
	echo building \"$(APP)\" into \"$(GOBIN)\" ... && \
	GOOS=linux $(GO) build -o "$(GOBIN)/$(APP)" cmd/pacman/*.go
	echo -=OK=-
