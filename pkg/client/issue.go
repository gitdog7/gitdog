package client

import (
	"context"
	"github.com/google/go-github/github"
	"log"
)

func (e *GitHubClient) FetchIssues() ([]*github.Issue, error) {
	result := make([]*github.Issue, 0)

	// list all issues for a given repo
	listOption := &github.IssueListByRepoOptions{
		State: "all",
	}
	listOption.PerPage = 100
	listOption.Page = 1
	totalIssues := 0

	for {
		issues, _, err := e.Client.Issues.ListByRepo(context.Background(),
			e.Owner, e.Repository, listOption)
		if err != nil {
			return result, err
		}

		log.Printf("Fetch %v issues, in page: %v", len(issues), listOption.Page)
		result = append(result, issues...)

		if len(issues) == 0 {
			break
		}

		totalIssues = totalIssues + len(issues)
		listOption.Page = listOption.Page + 1
	}

	log.Printf("Fetch %v issues in total.\n", totalIssues)
	return result, nil
}
