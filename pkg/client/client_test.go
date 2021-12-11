package client

import (
	"testing"
)

func Test_GitHubClient(t *testing.T) {

	client := NewGitHubClient("kubesphere", "console", "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")

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
