PACKAGE     = aws-power-toggle
METAPKG     = github.com/gbolo/aws-power-toggle/backend
DATE       ?= $(shell date +%FT%T%z)
VERSION     = 3.1
COMMIT_SHA ?= $(shell git rev-parse --short HEAD)
PKGS        = $(or $(PKG),$(shell $(GO) list ./...))
TESTPKGS    = $(shell $(GO) list -f '{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' $(PKGS))
BIN         = $(CURDIR)/bin
FRONTEND    = $(CURDIR)/frontend

GO      = go
TIMEOUT = 15

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

# Build

.PHONY: all
all: fmt dep $(BIN) frontend ; $(info $(M) building executable...) @ ## Build both backend and frontend
	$Q $(GO) build \
		-ldflags '-X $(METAPKG).Version=$(VERSION) -X $(METAPKG).BuildDate=$(DATE) -X $(METAPKG).CommitSHA=$(COMMIT_SHA)' \
		-o $(BIN)/$(PACKAGE)

.PHONY: backend
backend: ; $(info $(M) building backend executable...) @ ## Build backend binary
	$Q $(GO) build \
		-ldflags '-X $(METAPKG).Version=$(VERSION) -X $(METAPKG).BuildDate=$(DATE) -X $(METAPKG).CommitSHA=$(COMMIT_SHA)' \
		-o $(BIN)/$(PACKAGE)

.PHONY: docker
docker: clean ; $(info $(M) building docker image...)	@ ## Build docker image
	$Q docker build -t gbolo/$(PACKAGE):$(VERSION) .

.PHONY: frontend
frontend: ; $(info $(M) building web frontend ui...)	@ ## Build frontend
	$Q npm install --prefix $(FRONTEND) && npm run build --prefix $(FRONTEND)

# Tools

$(BIN):
	@mkdir -p $@
$(BIN)/%: | $(BIN) ; $(info $(M) building $(REPOSITORY)...)
	$Q tmp=$$(mktemp -d); \
	   env GO111MODULE=off GOCACHE=off GOPATH=$$tmp GOBIN=$(BIN) $(GO) get $(REPOSITORY) \
		|| ret=$$?; \
	   rm -rf $$tmp ; exit $$ret

DEP = $(BIN)/dep
$(BIN)/dep: REPOSITORY=github.com/golang/dep/cmd/dep

GOIMPORTS = $(BIN)/goimports
$(BIN)/goimports: REPOSITORY=golang.org/x/tools/cmd/goimports

GOLINT = $(BIN)/golint
$(BIN)/golint: REPOSITORY=golang.org/x/lint/golint

.PHONY: dep
dep: | $(DEP) ; $(info $(M) running dep...) @ ## Run dep ensure to fetch dependencies
	$Q $(DEP) ensure -v

.PHONY: fmt
fmt: ; $(info $(M) running gofmt...) @ ## Run gofmt on all source files
	$Q $(GO) fmt ./...

.PHONY: goimports
goimports: | $(GOIMPORTS) ; $(info $(M) running goimports...) @ ## Run goimports on backend source files
	$Q $(GOIMPORTS) -w backend/

.PHONY: lint
lint: | $(GOLINT) ; $(info $(M) running golint...) @ ## Run golint
	$Q $(GOLINT) -set_exit_status backend/

# Unit Test

.PHONY: test
test: ; $(info $(M) running go test...) @ ## Run go unit tests
	$Q $(GO) test -v -cover $(TESTPKGS)

# Misc

.PHONY: clean
clean: ; $(info $(M) cleaning...)	@ ## Cleanup everything
	@rm -rf $(BIN) $(CURDIR)/vendor $(FRONTEND)/dist $(FRONTEND)/node_modules

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo $(VERSION)

# credit: https://github.com/vincentbernat/hellogopher/blob/master/Makefile
