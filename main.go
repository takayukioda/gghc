package main

import (
	"os"
	"fmt"
	// "github.com/google/go-github/github"
)

func main() {
	var cmd string
	var opt []string
	var err error
	var args []string = os.Args

	if cmd, opt, err = parseArgs(args); err != nil {
		fmt.Fprintln(os.Stderr, "Error: %v", err)
	}
	fmt.Println("Command: %s", cmd)
	fmt.Println("Options: %s", opt)
}

func parseArgs(args []string) (string, []string, error) {
	subcmd, opt := args[1], args[2:]
	return subcmd, opt, nil
}
