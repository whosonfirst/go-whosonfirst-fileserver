prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	rmdeps deps fmt bin

deps:	self

fmt:
	go fmt cmd/*.go

bin: 	self
	@GOPATH=$(shell pwd) go build -o bin/wof-fileserver cmd/wof-fileserver.go
