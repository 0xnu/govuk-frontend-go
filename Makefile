.PHONY: test lint tidy fmt example

test:
	go test ./...

tidy:
	go mod tidy

fmt:
	gofmt -w $$(find . -name '*.go' -type f)

lint:
	golangci-lint run

example:
	go run cmd/example-gin-app/main.go