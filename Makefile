SHELL = bash

build:
	go build

clean:
	rm -rf pimacs
	go clean -testcache

fmt:
	gofmt -s -w -l .

checkfmt:
	test -z "$$(gofmt -l .)"

run: build
	./pimacs

test:
	PIMACS_TESTING=true go test -v ./...

lint:
	docker run -t --rm -v $$(pwd):/app -w /app golangci/golangci-lint:latest golangci-lint run -v

debug:
	dlv debug -- --load test.el
