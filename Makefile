build:
	@go build -o bin/rest-server

run: build
	@./bin/rest-server

test:
	@go test -v ./...
