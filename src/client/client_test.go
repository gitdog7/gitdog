package client

import (
	"testing"
)

func Test_GitHubClient(t *testing.T) {

	client := NewGitHubClient("kubesphere", "console", "")

	contributors, _ := client.FetchContributors()
	for _, c := range contributors {
		if *c.Contributions <= 0 {
			t.Error()
		}
	}

	if len(contributors) <= 0 {
		t.Error()
	}

	//issues, _ := client.FetchIssues()
	//print(len(issues))
}
