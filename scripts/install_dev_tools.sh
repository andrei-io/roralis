#!/bin/bash

cd ~
go install github.com/pressly/goose/v3/cmd/goose@latest
go install github.com/go-swagger/go-swagger/cmd/swagger@latest
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0