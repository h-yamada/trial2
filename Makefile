.PHONY: test

NAME := trial
VERSION := test
#VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
SRCS = $(shell git ls-files '*.go')

test:
	go test -cover -v
