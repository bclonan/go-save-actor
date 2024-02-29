package githubapi

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var client *github.Client

func Init(token string) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	client = github.NewClient(tc)
}

func GetRepo(owner string, repo string) (*github.Repository, error) {
	repository, _, err := client.Repositories.Get(context.Background(), owner, repo)
	return repository, err
}

func CreateIssue(owner string, repo string, issue *github.IssueRequest) (*github.Issue, error) {
	newIssue, _, err := client.Issues.Create(context.Background(), owner, repo, issue)
	return newIssue, err
}