SHELL = bash
GOLANGCI_LINT_CACHE = ~/.cache/golangci-lint/latest

build:
	go build

clean-test:
	go clean -testcache

clean: clean-test
	rm -rf pimacs
	sudo rm -rf $(GOLANGCI_LINT_CACHE)

fmt:
	gofmt -s -w -l .

checkfmt:
	test -z "$$(gofmt -l .)"

run: build
	./pimacs

test:
	go test -v ./...

lint:
	docker run -t --rm -v $$(pwd):/app -v $(GOLANGCI_LINT_CACHE):/root/.cache -w /app golangci/golangci-lint:latest golangci-lint run -v

pre-push: fmt lint test
