#!/bin/sh

cd /go/src/app
go install -v ./..
go test -v ./test