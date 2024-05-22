package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func pipe_commands() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	var echoStderr bytes.Buffer
	var wcStdout bytes.Buffer
	var wcStderr bytes.Buffer
	echoCmd := exec.CommandContext(ctx, "echo", "-sd", "hello")
	echoCmd.Stderr = &echoStderr
	wcCmd := exec.CommandContext(ctx, "wc", "-l")
	wcCmd.Stdin, _ = echoCmd.StdoutPipe()
	wcCmd.Stdout = &wcStdout
	wcCmd.Stderr = &wcStderr
	if err := wcCmd.Start(); err != nil {
		log.Fatalf("starting wc process: %s", err)
	}
	if err := echoCmd.Run(); err != nil {

		log.Fatalf("running echo process: %s", echoCmd.String())
	}
	if err := wcCmd.Wait(); err != nil {
		fmt.Printf("stderr: %s\n", wcStderr.String())
		log.Fatal("waiting for wc")
	}
	if count, err := strconv.Atoi(strings.TrimSpace(wcStdout.String())); err != nil {
		log.Fatalf("converting to int: %s", err)
	} else {
		fmt.Printf("count: %d\n", count)
	}
	fmt.Printf("stderr: %s\n", wcStderr.String())
}
