package repo

import (
	"github.com/gitdog7/gitdog/src/client"
	"log"
)

func BuildGitHubRepo(client *client.GitHubClient) *GitHubRepo {
	repo := &GitHubRepo{Owner: client.Owner, Repository: client.Repository}
	repo.UpdateRepo(client)
	return repo
}

// update repo
func (repo *GitHubRepo) UpdateRepo(client *client.GitHubClient) error {

	// fetch contributors
	contributors, err := client.FetchContributors()
	if err != nil {
		log.Fatalln("failed to fetch contributors.")
	}
	repo.SetContributors(contributors)

	// fetch followings for each contributor
	followingCnt := 0
	for idx, c := range contributors {
		followings, _ := client.FetchFollowings(c)
		repo.AppendFollowings(c, followings)
		followingCnt = followingCnt + len(followings)
		log.Printf("### Fetch followings for all contributors, progress: %v/%v", idx, len(contributors))
	}

	// fetch members
	members, err := client.FetchMembers()
	repo.SetMembers(members)

	log.Printf("update repository data successfully.")
	log.Printf("total %v contributors.", len(contributors))
	log.Printf("total %v members.", len(members))
	log.Printf("total %v following relationships.", followingCnt)

	// fetch issues
	issues, err := client.FetchIssues()
	if err != nil {
		return err
	}

	repo.Issues = issues

	//for _, issue := range issues {
	//	log.Printf("%v %v %v", *issue.User, *issue.ClosedAt, issue.IsPullRequest())
	//}

	return nil
}
