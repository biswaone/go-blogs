build:
	@go build -o bin/goblogs

run: build
	@./bin/goblogs

test:
	@go test -v ./...