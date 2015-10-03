.PHONY: setup test

default: test

setup:
	go get -t ./...
	go get github.com/smartystreets/goconvey
	go get github.com/codegangsta/gin

test:
	go test -timeout=60s ./...
