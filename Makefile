.PHONY: setup test run build install

default: run

setup:
	go get -t ./...
	go get -u github.com/mailgun/godebug
	go get -u github.com/jteeuwen/go-bindata/...
	go get -u github.com/cespare/reflex
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install --update

build:
	go-bindata templates/
	go build

run:
	make build
	./etch

install:
	make build
	sudo mv ./etch /usr/local/bin

test:
	go test -timeout=60s ./...
