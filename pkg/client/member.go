package client

import (
	"context"
	"github.com/google/go-github/github"
	"log"
)

func (e *GitHubClient) FetchMembers() ([]*github.User, error) {
	result := make([]*github.User, 0)

	// try to get the repository
	repo, _, err := e.Client.Repositories.Get(context.Background(), e.Owner, e.Repository)
	if err != nil {
		return result, err
	}

	// get organization
	org := *repo.Organization.Login
	option := &github.ListMembersOptions{}
	option.PerPage = 100
	option.Page = 1
	option.PublicOnly = true

	for {
		users, _, err := e.Client.Organizations.ListMembers(context.Background(), org, option)
		if err != nil {
			return result, err
		}

		if len(users) == 0 {
			break
		}

		result = append(result, users...)
		option.Page = option.Page + 1
	}

	log.Printf("### Total %v members fetched.\n", len(result))
	return result, nil
}
