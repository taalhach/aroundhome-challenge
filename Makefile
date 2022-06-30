clean:
	@rm -rf ./bin

build: clean
	@mkdir -p bin
	@go build -o bin/aroundhome cmd/aroundhome/main.go

test-unit:
	@go test -race ./...

run:
	@go run cmd/aroundhome/main.go serve_api
