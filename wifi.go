package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func before(name string, args []string) {
	fmt.Println(">", name, strings.Join(args, " "))
}

func command(name string, args ...string) {
	before(name, args)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err, ok := err.(*exec.ExitError); ok {
		os.Exit(err.ExitCode())
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	command("netsh", "wlan", "stop", "hostednetwork")
	command("netsh", "wlan", "start", "hostednetwork")
}
