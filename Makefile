clean:
	@rm -rf ./bin

build: clean
	@mkdir -p bin
	@go build -o bin/challenge cmd/challenge/main.go

test-unit:
	@go test -race ./...

run:
	@go run cmd/challenge/main.go serve_api
