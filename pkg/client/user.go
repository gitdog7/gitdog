package client

import (
	"context"
	"github.com/google/go-github/github"
	"log"
)

func (e *GitHubClient) FetchFollowers(c *github.Contributor) ([]*github.User, error) {
	result := make([]*github.User, 0)
	listOption := &github.ListOptions{}
	listOption.Page = 1
	listOption.PerPage = 100

	for {
		followers, _, err := e.Client.Users.ListFollowers(context.Background(), *c.Login, listOption)
		if err != nil {
			return result, err
		}

		if len(followers) == 0 {
			break
		}

		result = append(result, followers...)
		listOption.Page = listOption.Page + 1
	}

	log.Printf("### Total %v followers of user %v .\n", len(result), *c.Login)
	return result, nil
}

func (e *GitHubClient) FetchFollowings(c *github.Contributor) ([]*github.User, error) {
	result := make([]*github.User, 0)
	listOption := &github.ListOptions{}
	listOption.Page = 1
	listOption.PerPage = 100

	for {
		followers, _, err := e.Client.Users.ListFollowing(context.Background(), *c.Login, listOption)
		if err != nil {
			return result, err
		}

		if len(followers) == 0 {
			break
		}

		result = append(result, followers...)
		listOption.Page = listOption.Page + 1
	}

	log.Printf("### Total %v followings of user %v .\n", len(result), *c.Login)
	return result, nil
}
