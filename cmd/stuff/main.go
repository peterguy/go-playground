package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestX(t *testing.T) {
	const a = 257
	// From github.com/stretchr/testify/require
	require.Greater(t, int8(10), a)
}

func stuff() {
	stuff := `
hello
world
`
	y := strings.ReplaceAll(stuff, "\n", "\n\t")
	fmt.Println(y)
}
