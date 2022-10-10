test:
	ginkgo -r
lint:
	find src -type f -name *.go -exec gofmt -s -w {} \;
	find src -type f -name *.go -exec goimports -w {} \;
	golangci-lint run -c .golangci.yml
run:
	go run src/main.go
