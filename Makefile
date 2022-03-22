test:
	go test -cover -v -tags integration ./...

lint: ## Lint the files local env
	golangci-lint run --timeout=5m -c .golangci.yml
