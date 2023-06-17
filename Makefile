PACKAGE_NAME := "$(shell head -n 1 go.mod | cut -d ' ' -f2)"

dep:
	@go mod tidy

lint:
	@revive -config ./revive.toml ./...

test-coverage:
	@go test -race -coverprofile=coverage.txt -covermode=atomic ./...

build:
	go build -o myapp ./cmd/myapp

.PHONY: lint test-coverage build
