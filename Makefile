#!/bin/bash

fmt:
	goimports -l -w ./src/

dev: fmt
	go build -o output/bin/taurus-MacOs ./src