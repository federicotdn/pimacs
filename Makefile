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

debug:
	dlv debug -- --load test.el
