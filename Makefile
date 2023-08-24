SHELL = bash
PIMACS_LISP = $(CURDIR)/lisp


build:
	go build

clean:
	rm -rf pimacs
	go clean -testcache

fmt:
	gofmt -s -w -l .

run: build
	./pimacs

test:
	PIMACS_TESTING=true PIMACS_LISP=$(PIMACS_LISP) go test -v ./...

debug:
	dlv debug -- --load test.el
