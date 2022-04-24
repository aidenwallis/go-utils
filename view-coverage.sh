#!/bin/sh

# This small script simply runs the tests and opens a browser window of the results.
# From: https://stackoverflow.com/a/27284510

t="/tmp/go-cover.$$.tmp"
go test -coverprofile=$t -race ./... $@ && go tool cover -html=$t && unlink $t
