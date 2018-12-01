PACKAGE     = aws-power-toggle
DATE       ?= $(shell date +%FT%T%z)
VERSION     = 3.0
COMMIT_SHA ?= $(shell git rev-parse --short HEAD)
PKGS        = $(or $(PKG),$(shell $(GO) list ./...))
TESTPKGS    = $(shell $(GO) list -f '{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' $(PKGS))
BIN         = $(CURDIR)/bin

GO      = go
TIMEOUT = 15

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: all
all: fmt dep $(BIN) ; $(info $(M) building executable...) @ ## Build program binary
	$Q $(GO) build \
		-ldflags '-X main.Version=$(VERSION) -X main.BuildDate=$(DATE) -X main.CommitSHA=$(COMMIT_SHA)' \
		-o $(BIN)/$(PACKAGE)

.PHONY: docker
docker: clean ; $(info $(M) building docker image...)	@ ## Build docker imaage
	$Q docker build -t gbolo/$(PACKAGE):$(VERSION) .

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

.PHONY: dep
dep: | $(DEP) ; $(info $(M) running dep...) @ ## Run dep ensure to fetch dependencies
	$Q $(DEP) ensure -v

.PHONY: fmt
fmt: ; $(info $(M) running gofmt...) @ ## Run gofmt on all source files
	$Q $(GO) fmt ./...

.PHONY: test
test: ; $(info $(M) running go test...) @ ## Run go test
	$Q $(GO) test -v -cover $(TESTPKGS)

# Misc

.PHONY: clean
clean: ; $(info $(M) cleaning...)	@ ## Cleanup everything
	@rm -rf $(BIN)

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo $(VERSION)

# credit: https://github.com/vincentbernat/hellogopher/blob/master/Makefile
