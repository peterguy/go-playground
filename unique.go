package main

import (
	"fmt"
	"sort"
	"strings"
)

func unique() {
	x := []string{"z", "a", "c", "a", "b", "c", "c", "d", "e", "f", "g", "h", "h", "h"}
	sort.Strings(x)
	u := 0
	for i := 1; i < len(x); i++ {
		if x[u] != x[i] {
			if i > u+1 {
				x[u+1] = x[i]
			}
			u++
		}
	}
	fmt.Printf("%s\n", strings.Join(x[:u+1], ", "))
}
