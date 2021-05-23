source_files := $(wildcard */**.go)
packages := `go list ./...`

build:
	@go build $(packages)

fmt:
	@goimports -w ${source_files}

install:
	@go mod download

install-tools:
	@curl https://language-version.byted.org/bytedance-golangci-lint/darwin/amd64/v1.0.0/golangci-lint -o $(go env GOPATH)/bin/golangci-lint && chmod +x $(go env GOPATH)/bin/golangci-lint
	@cd ~; go get golang.org/x/tools/cmd/goimports

lint:
	@$(shell go env GOPATH)/bin/golangci-lint run --exclude kitex_gen --color always --concurrency 4 --verbose --sort-results

ci: lint test build

test:
	@go test -v -coverprofile=coverage.out $(packages) -run=$(name)
	@go tool cover -func=coverage.out

clean:
	@-go clean -cache
	@-rm -f coverage.out
