.PHONY: setup test

default: test

setup:
	go get -t ./...
	go get -u github.com/mailgun/godebug
	go get -u github.com/jteeuwen/go-bindata/...
	go get -u github.com/cespare/reflex
	go get -u github.com/alecthomas/gometalinter

test:
	go test -timeout=60s ./...
