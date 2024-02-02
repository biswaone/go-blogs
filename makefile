build:
	@go build -o bin/go-blogs

run: build
	@./bin/go-blogs

test:
	@go test -v ./...