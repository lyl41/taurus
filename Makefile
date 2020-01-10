#!/bin/bash

fmt:
	goimports -l -w ./src/

dev: fmt
	go build -o output/bin/taurus-MacOs ./src

linux: fmt
	GOOS=linux GOARCH=amd64 go build -o output/bin/taurus-server ./src
