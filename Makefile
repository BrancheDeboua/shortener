.DEFAULT_GOAL := run

run:
	@go run main.go

vet: fmt
	@go vet ./...

fmt:
	@go fmt ./...

test: vet
	@go test ./...

