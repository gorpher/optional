#! /bin/bash -x
basepath=$(pwd)
echo "now at ${basepath}"
goimports -w  .
golangci-lint run  -c .golangci.yml
go mod tidy
