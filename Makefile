SHELL = bash
PIMACS_LISP = $(CURDIR)/lisp


build: clean
	go build

clean:
	rm -rf pimacs
	go clean -testcache

fmt:
	gofmt -s -w -l .

run: build
	./pimacs

test: clean
	PIMACS_TESTING=true PIMACS_LISP=$(PIMACS_LISP) go test -v ./...

debug:
	dlv debug -- --load test.el

code-tags:
	grep -n -r --exclude Makefile '// TAGS: ' | sed 's/\/\/ TAGS: //' | column -t
