#!/usr/bin/env bash
rm -Rf ./dist/
go build -o ./dist/gozos && cd dist && ./gozos && cd ..