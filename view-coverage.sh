#!/bin/sh

# This small script simply runs the tests and opens a browser window of the results.
# From: 

t="/tmp/go-cover.$$.tmp"
go test -coverprofile=$t -race ./... $@ && go tool cover -html=$t && unlink $t
