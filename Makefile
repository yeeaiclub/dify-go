.PHONY: precommit
precommit:
precommit: all

.PHONY: all
all:
all: fmt lint test clean

.PHONY:	lint
lint:
	@echo "lint code..."
	@golangci-lint run -c .golangci.yaml

.PHONY:	fmt
fmt:
	@echo "format code..."
	@goimports -l -w $$(find . -type f -name '*.go'  -not -path "./.idea/*" -not -name '*.pb.go' -not -name '*mock*.go')
	@gofumpt -w .

.PHONY: test
test:
	@echo "Running tests..."
	@go test -v -race -timeout=30s -coverprofile=coverage.out ./...

.PHONY: clean
clean:
	rm -f coverage.*