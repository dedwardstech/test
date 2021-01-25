test:
	go test --race ./...

fmt:
	gofmt -w ./**/*.go