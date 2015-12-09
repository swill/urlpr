#!/usr/bin/env bash

esc -o static.go static
gox -osarch="darwin/amd64" -output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}"