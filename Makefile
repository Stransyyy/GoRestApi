build: 
	@go build -o bin/GoRestApi

run: build
	@./bin/GoRestApi

test:
	@go test -v ./...