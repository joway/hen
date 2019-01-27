#!/usr/bin/env bash

go list ./... | xargs -L1 golint
gofmt -s -w  ./**/*.go
./dep.sh
