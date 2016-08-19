.PHONY: setup test run build install

default: run

setup:
	go get -t ./...
	go get -u github.com/jteeuwen/go-bindata/...
	go get -u github.com/alecthomas/gometalinter
	go get -u github.com/cortesi/modd/cmd/modd
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
