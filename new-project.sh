#!/usr/bin/env bash

project=${1}

[ -e "cmd/${project}" ] && {
    echo "cmd/${project} already exists!"
    exit 1
}

mkdir -p "cmd/${project}"

cat > "cmd/${project}/main.go" <<EOF
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello!")
    fmt.Println("cmd/${project}/main.go is all ready to go now!")
}
EOF

go run "./cmd/${project}"
