SHELL = bash

clean:
	rm -rf pimacs

build:
	go build

run:
	go build -gcflags=all="-N -l"
	./pimacs

test:
	go test -v ./...
