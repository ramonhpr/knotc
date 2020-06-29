#!/usr/bin/env bash

version=`cat VERSION`
output_name=knotc

go build -ldflags="-X main.version=$version" -o $output_name cmd/main.go