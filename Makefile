build:
	@go build -o untitled-project cmd/main.go

test:
	@go test -v ./...

run: build
	@./untitled-project