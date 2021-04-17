#!/bin/bash

go mod tidy
GOOS=windows GOARCH=amd64 go build
