test:
	ginkgo -r
lint:
	find src -type f -name *.go -exec gofmt -s -w {} \;
	find src -type f -name *.go -exec goimports -w {} \;
	golangci-lint run -c .golangci.yml
image:
	docker build -t freecocoa:latest .
run:
	go run src/main.go
run_docker:
	docker run -p 7123:7123 freecocoa
