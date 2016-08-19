.PHONY: setup test

default: test

setup:
	go get -t ./...
	go get -u github.com/cortesi/modd/cmd/modd
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install --update

test:
	go test -timeout=60s ./...
