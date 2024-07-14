#!/usr/bin/env zsh

go test $(go list ./... | grep -v /features/)

