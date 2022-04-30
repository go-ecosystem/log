SHELL := /bin/bash
BASEDIR = $(shell pwd)

all: fmt imports mod lint test
first:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
fmt:
	gofmt -w .
imports:
	goimports -w .
mod:
	go mod tidy
lint:
	golangci-lint run -c .golangci.yml
.PHONY: test
test: mod
	go test -gcflags=-l -coverpkg=./... -coverprofile=coverage.data ./...
	go tool cover -html=coverage.data -o coverage.html;
update:
	# https://github.com/golang/go/wiki/Modules#how-to-upgrade-and-downgrade-dependencies
	go list -u -f '{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}' -m all 2> /dev/null	
help:
	@echo "fmt - format the source code"
	@echo "imports - goimports"
	@echo "mod - go mod tidy"
	@echo "lint - run golangci-lint"
	@echo "test - unit test"
	@echo "update - list updateable packages"