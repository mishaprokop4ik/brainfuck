GOLINT := golangci-lint
test:
	go test -cover ./...
lint:
	$(GOLINT) run --timeout=5m -c .golangci.yml
fmt:
	go fmt ./...
precommit: fmt lint test
