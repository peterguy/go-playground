#!/usr/bin/env bash

for dir in cmd/*/; do
    name=$(basename "$dir")
    go build -o "bin/$name" "./cmd/$name"
done
