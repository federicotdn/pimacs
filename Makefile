SHELL = bash
GOLANGCI_LINT_CACHE = ~/.cache/golangci-lint/latest

build:
	go build

clean:
	rm -rf pimacs
	go clean -testcache
	rm -rf $(GOLANGCI_LINT_CACHE)

fmt:
	gofmt -s -w -l .

checkfmt:
	test -z "$$(gofmt -l .)"

run: build
	./pimacs

test:
	PIMACS_TESTING=true go test -v ./...

lint:
	docker run -t --rm -v $$(pwd):/app -v $(GOLANGCI_LINT_CACHE):/root/.cache -w /app golangci/golangci-lint:latest golangci-lint run -v

debug:
	dlv debug -- --load test.el
