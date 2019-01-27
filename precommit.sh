#!/usr/bin/env bash

gofmt -s -w  ./**/*.go
./dep.sh
