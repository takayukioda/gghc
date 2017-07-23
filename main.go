package main

import (
	"flag"
	"fmt"
	"os"
	// "github.com/golang/x/oauth2"
	// "github.com/google/go-github/github"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "this command requires at least one argument")
		os.Exit(1)
	}

	user := flag.String("user", "", "GitHub username")
	repo := flag.String("repo", "", "Repository name")
	token := flag.String("token", "", "Access token")

	flag.Parse()

	args := flag.Args()

	target := args[0]

	switch target {
	default:
		fmt.Println("Option[user]:", *user)
		fmt.Println("Option[repo]:", *repo)
		fmt.Println("Option[token]:", *token)
		fmt.Println("Target:", target)
	}
}
