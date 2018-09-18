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
	# @GOPATH=$(GOPATH) go get -u "github.com/vaughan0/go-ini"

bin:	self fmt
	@GOPATH=$(GOPATH) go build -o bin/wof-fileserver cmd/wof-fileserver.go

fmt:
	go fmt cmd/*.go
