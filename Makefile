CWD=$(shell pwd)
VENDORGOPATH := $(CWD)/vendor:$(CWD)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:	prep

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	rmdeps deps bin

deps:
	@GOPATH=$(GOPATH) go get -u "github.com/vaughan0/go-ini"
	@GOPATH=$(GOPATH) go get -u "golang.org/x/net/html"
	@GOPATH=$(GOPATH) go get -u "golang.org/x/oauth2"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-httpony"

bin:	self fmt
	@GOPATH=$(GOPATH) go build -o bin/wof-fileserver cmd/wof-fileserver.go

fmt:
	go fmt cmd/*.go
