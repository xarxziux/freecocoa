test:
	ginkgo -r
lint:
	find src -type f -name *.go -exec gofmt -s -w {} \;
	find src -type f -name *.go -exec goimports -w {} \;
	golangci-lint run --enable stylecheck,unparam,gocritic,whitespace
run:
	go run src/main.go
