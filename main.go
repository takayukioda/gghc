package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	user := flag.String("user", "", "GitHub username")
	repo := flag.String("repo", "", "Repository name")
	token := flag.String("token", "", "Access token")

	flag.Parse()

	// configuring github token
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	args := flag.Args()

	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <resource> <action>\n", os.Args[0])
		os.Exit(1)
	}

	target := args[0]
	action := args[1]
	switch target {
	case "labels":
		if action != "list" {
			fmt.Fprintln(os.Stderr, "Undefined action:", action)
			os.Exit(1)
		}
		labels, _, err := client.Issues.ListLabels(ctx, *user, *repo, nil)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, l := range labels {
			fmt.Println("Label:", l.String())
		}
	default:
		fmt.Println("Option[user]:", *user)
		fmt.Println("Option[repo]:", *repo)
		fmt.Println("Option[token]:", *token)
		fmt.Println("Target:", target)
	}

	os.Exit(0)
}
