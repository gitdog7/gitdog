package client

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
)

type GitHubClient struct {
	Client     *github.Client
	Owner      string
	Repository string
}

func NewGitHubClient(owner string, repository string, token string) *GitHubClient {
	ret := &GitHubClient{}
	ret.Owner = owner
	ret.Repository = repository

	if token != "" {
		// create token
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)

		// create http client with oauth
		tc := oauth2.NewClient(context.Background(), ts)

		// create client
		ret.Client = github.NewClient(tc)
	} else {
		// create client without token, use default rate limit
		ret.Client = github.NewClient(nil)
	}

	return ret
}

func (c *GitHubClient) IsClientValid() bool {
	user, _, err := c.Client.Users.Get(context.Background(), "live77")
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return *user.Login == "live77"
}

// TODO: FIXME
func (c *GitHubClient) IsRepoExist() bool {
	return true
}
