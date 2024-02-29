package adapter

import (
	"github.com/google/go-github/github"
	"context"
)

type GithubAPIAdapter struct {
	client *github.Client
}

func NewGithubAPIAdapter() *GithubAPIAdapter {
	client := github.NewClient(nil)
	return &GithubAPIAdapter{client: client}
}

func (g *GithubAPIAdapter) GetRepositories(user string) ([]*github.Repository, error) {
	ctx := context.Background()
	repos, _, err := g.client.Repositories.List(ctx, user, nil)
	return repos, err
}

func (g *GithubAPIAdapter) CreateIssue(owner string, repo string, issue *github.IssueRequest) (*github.Issue, error) {
	ctx := context.Background()
	newIssue, _, err := g.client.Issues.Create(ctx, owner, repo, issue)
	return newIssue, err
}