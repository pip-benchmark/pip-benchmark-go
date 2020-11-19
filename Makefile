.PHONY: all build clean install uninstall fmt simplify check run test

install:
	@go install ./app/main.go

run: install
	@go run ./app/main.go

test:
	@go clean -testcache && go test -v ./test/...

plugins:
	@go build -buildmode=plugin -o ./plugin/standardbenchmarks.so ./plugin/main.go