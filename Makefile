# Variables -------------------------------------------------------------------------------------------------------------

PACKAGE     = aws-power-toggle
METAPKG     = github.com/gbolo/aws-power-toggle/backend
DATE       ?= $(shell date +%FT%T%z)
VERSION     = 3.3
COMMIT_SHA ?= $(shell git rev-parse --short HEAD)
LDFLAGS     = -X $(METAPKG).Version=$(VERSION) -X $(METAPKG).BuildDate=$(DATE) -X $(METAPKG).CommitSHA=$(COMMIT_SHA)
PKGS        = $(or $(PKG),$(shell $(GO) list ./...))
TESTPKGS    = $(shell $(GO) list -f '{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' $(PKGS))
BIN         = $(CURDIR)/bin
FRONTEND    = $(CURDIR)/frontend
GO          = go

V ?= 0
Q  = $(if $(filter 1,$V),,@)
M  = $(shell printf "\033[34;1m▶\033[0m")

# Targets to build Go tools --------------------------------------------------------------------------------------------

$(BIN):
	@mkdir -p $@

$(BIN)/%: | $(BIN) ; $(info $(M) building $(REPOSITORY)...)
	$Q tmp=$$(mktemp -d); \
	   env GOPATH=$$tmp GOBIN=$(BIN) $(GO) get $(REPOSITORY) \
		|| ret=$$?; \
	   rm -rf $$tmp ; exit $$ret

GOIMPORTS = $(BIN)/goimports
$(BIN)/goimports: REPOSITORY=golang.org/x/tools/cmd/goimports

GOLINT = $(BIN)/golint
$(BIN)/golint: REPOSITORY=golang.org/x/lint/golint

# Targets for our app --------------------------------------------------------------------------------------------------

.PHONY: all
all: $(BIN) frontend backend;                                 @ ## Build both backend and frontend

.PHONY: backend
backend: ; $(info $(M) building backend executable...)        @ ## Build backend binary
	$Q $(GO) build -ldflags '$(LDFLAGS)' -o $(BIN)/$(PACKAGE)

.PHONY: docker
docker: clean ; $(info $(M) building docker image...)	      @ ## Build docker image
	$Q docker build -t gbolo/$(PACKAGE):$(VERSION) .

.PHONY: frontend
frontend: ; $(info $(M) building web frontend ui...)	      @ ## Build frontend
	$Q npm install --prefix $(FRONTEND) \
	   && npm run build --prefix $(FRONTEND)

.PHONY: fmt
fmt: ; $(info $(M) running gofmt...)                          @ ## Run gofmt on all source files
	$Q $(GO) fmt ./...

.PHONY: goimports
goimports: | $(GOIMPORTS) ; $(info $(M) running goimports...) @ ## Run goimports on backend source files
	$Q $(GOIMPORTS) -w backend/

.PHONY: lint
lint: | $(GOLINT) ; $(info $(M) running golint...)            @ ## Run golint
	$Q $(GOLINT) -set_exit_status backend/

.PHONY: test
test: ; $(info $(M) running go test...)                       @ ## Run go unit tests
	$Q $(GO) test -v -cover $(TESTPKGS)

.PHONY: clean
clean: ; $(info $(M) cleaning...)                             @ ## Cleanup everything
	@rm -rf $(BIN) $(CURDIR)/vendor $(FRONTEND)/dist $(FRONTEND)/node_modules

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo $(VERSION)

# credit: https://github.com/vincentbernat/hellogopher/blob/master/Makefile
