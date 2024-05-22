package main

import (
	"fmt"
	"os/exec"
)

func doit(stuff []string) {
	cmd := exec.Command(stuff[0], stuff[1:]...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running command:", err)
		return
	}
	fmt.Println(string(output))
}

func execute() {
	doit([]string{"/bin/sh", "-c", "echo", "separated"})
	doit([]string{"/bin/sh -c \"echo all together\""})
}
