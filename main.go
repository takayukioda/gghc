package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	// "github.com/k0kubun/pp"
	"golang.org/x/oauth2"
)

const (
	DEFAULT_PERPAGE = 30
)

func main() {
	token := os.Getenv("GGHC_GITHUB_TOKEN")

	user := flag.String("user", "", "GitHub username")
	repo := flag.String("repo", "", "Repository name")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <resource> <action>\n", os.Args[0])
		os.Exit(1)
	}
	target := args[0]
	action := args[1]

	ctx := context.Background()
	client := newGitHubClient(ctx, token)

	switch target {
	case "labels":
		labels(ctx, client, *user, *repo, action)
	default:
		fmt.Println("Option[user]:", *user)
		fmt.Println("Option[repo]:", *repo)
		fmt.Println("Target:", target)
		fmt.Println("Action:", action)
	}

	os.Exit(0)
}

func newGitHubClient(ctx context.Context, token string) *github.Client {
	// configuring github client by personal access token
	// reference: <https://github.com/google/go-github#authentication>
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

func labels(ctx context.Context, client *github.Client, user string, repo string, action string) {
	if action != "list" {
		fmt.Fprintln(os.Stderr, "Unknow action:", action)
		os.Exit(1)
	}
	labels := make([]*github.Label, 0, DEFAULT_PERPAGE)
	opt := github.ListOptions{PerPage: DEFAULT_PERPAGE}

	for {
		ls, resp, err := client.Issues.ListLabels(ctx, user, repo, &opt)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		labels = append(labels, ls...)
		opt.Page = resp.NextPage
		fmt.Println("NextPage:", resp.NextPage)
		fmt.Println("LastPage:", resp.LastPage)
		if resp.NextPage == 0 {
			break
		}
	}

	fmt.Println("You've got", len(labels), "labels")
	for _, l := range labels {
		fmt.Println("Label:", *(l.Name))
	}
}
