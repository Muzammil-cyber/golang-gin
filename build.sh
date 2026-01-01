#!/usr/bin/env bash

set -xe

# install packages and dependencies
go get -u github.com/gin-gonic/gin
go get github.com/go-playground/validator/v10
go get github.com/tpkeeper/gin-dump

# build the application
GOOS=linux GOARCH=amd64 go build -o bin/application main.go

