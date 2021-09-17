SHELL = bash

clean:
	rm -rf pimacs

build:
	go build

fmt:
	gofmt -s -w -l .

run:
	go build && ./pimacs

run-file:
	go build && ./pimacs --load test.el

test:
	go test -v ./...

debug:
	dlv debug -- --load test.el

code-tags:
	grep -n -r --exclude Makefile '// TAGS: ' | sed 's/\/\/ TAGS: //' | column -t
