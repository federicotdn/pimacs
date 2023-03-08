SHELL = bash
PIMACS_ELISP = $(CURDIR)/elisp

clean:
	rm -rf pimacs

build:
	go clean
	go build

fmt:
	gofmt -s -w -l .

run:
	make build && ./pimacs

run-file:
	make build && ./pimacs --load test.el

test:
	go clean -testcache
	PIMACS_TESTING=true PIMACS_ELISP=$(PIMACS_ELISP) go test -v ./...

debug:
	dlv debug -- --load test.el

code-tags:
	grep -n -r --exclude Makefile '// TAGS: ' | sed 's/\/\/ TAGS: //' | column -t
