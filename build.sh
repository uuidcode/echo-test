#!/usr/bin/env bash
rm -rf dist
mkdir dist
go build -o dist/server server.go
cp -r template static dist