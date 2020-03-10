.PHONY: all build clean install uninstall fmt simplify check run test

install:
	@go install main.go

run: install
	@go run main.go

test:
	@go test -v ./test/...

plugin:
	@go build -buildmode=plugin -o ./plugin/standardbenchmarks.so ./plugin/main.go