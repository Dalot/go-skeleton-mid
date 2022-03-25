.PHONY: build test check

GO ?= go

build: 
	ENV=${ENV} CGO_ENABLED=0 $(GO) build -o ./server ./cmd/app

test: 
	$(GO) test -cover ./...

check:
	$(GO) vet $$($(GO) list ./...)
	$(GOLINT) $$($(GO) list ./...)
