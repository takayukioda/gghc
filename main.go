package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	// "github.com/k0kubun/pp"
	"golang.org/x/oauth2"
)

const (
	DefaultPerPage = 30
)

const (
	ExitOk = 0
	ExitError
)

var client *github.Client

func main() {
	os.Exit(run())
}

func run() int {
	token := os.Getenv("GGHC_GITHUB_TOKEN")

	user := flag.String("user", "", "GitHub username")
	repo := flag.String("repo", "", "Repository name")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <resource> <action>\n", os.Args[0])
		return ExitError
	}
	target := args[0]
	action := args[1]

	ctx := context.Background()
	client = newGitHubClient(ctx, token)

	switch target {
	case "labels":
		err := labels(ctx, *user, *repo, action)
		if err != nil {
			return ExitError
		}
	default:
		fmt.Println("Option[user]:", *user)
		fmt.Println("Option[repo]:", *repo)
		fmt.Println("Target:", target)
		fmt.Println("Action:", action)
	}
	return ExitOk
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

func labels(ctx context.Context, user string, repo string, action string) error {
	if action != "list" {
		fmt.Fprintln(os.Stderr, "Unknow action:", action)
		return errors.New("Unknow action given")
	}
	labels, err := getAllLabels(ctx, user, repo)
	if err != nil {
		return err
	}

	fmt.Println("You've got", len(labels), "labels")
	for _, l := range labels {
		fmt.Println("Label:", *(l.Name))
	}
	return nil
}

func getAllLabels(ctx context.Context, user string, repo string) ([]github.Label, error) {
	allp := make([]*github.Label, 0, DefaultPerPage)
	opt := github.ListOptions{PerPage: DefaultPerPage}
	for {
		labels, resp, err := client.Issues.ListLabels(ctx, user, repo, &opt)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil, err
		}
		allp = append(allp, labels...)
		opt.Page = resp.NextPage

		// When there's no next page, GitHub omits next page information <https://developer.github.com/v3/#pagination>
		// go-github resolves omitted information by setting with zero
		// <https://github.com/google/go-github/blob/35d38108ba83757e9b6fc00e4ba7e2d597c651be/github/github.go#L297-L314>
		if resp.NextPage == 0 {
			break
		}
	}

	labels := make([]github.Label, 0, len(allp))
	for _, label := range allp {
		labels = append(labels, *label)
	}
	return labels, nil
}
