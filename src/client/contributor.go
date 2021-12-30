package client

import (
	"context"
	"github.com/google/go-github/github"
	"log"
)

func (e *GitHubClient) FetchContributors() ([]*github.Contributor, error) {
	result := make([]*github.Contributor, 0)
	nextPage := 1

	for {
		listOption := &github.ListContributorsOptions{}
		listOption.Page = nextPage

		contributors, _, err := e.Client.Repositories.ListContributors(context.Background(),
			e.Owner, e.Repository, listOption)
		if err != nil {
			return result, err
		}

		//log.Printf("get %v contributors, in page: %v\r", len(contributors), listOption.Page)
		if len(contributors) == 0 {
			break
		}

		result = append(result, contributors...)
		nextPage = nextPage + 1
	}

	log.Printf("### Total %v contributors info fetched.\n", len(result))
	return result, nil
}
