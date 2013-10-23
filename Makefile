all:
	@mkdir -p bin/
	@go build -o bin/carpenter carpenter.go

get-deps:
	@go get -d -v ./...

test:
	@go test -v ./...

format:
	@go fmt ./...

clean:
	@rm -rf bin/
