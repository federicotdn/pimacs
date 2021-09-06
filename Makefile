SHELL = bash

clean:
	rm -rf pimacs

build:
	go build

fmt:
	gofmt -s -w -l .

run:
	go build -gcflags=all="-N -l"
	./pimacs

test:
	go test -v ./...
