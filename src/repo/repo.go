package repo

import (
	"encoding/json"
	"github.com/google/go-github/github"
	"io/ioutil"
	"log"
)

// GitHubRepo Contains all the data we need to analysis a given repository.
type GitHubRepo struct {
	Owner      string
	Repository string

	// All contributors, Login as primary key
	Contributors map[string]*github.Contributor

	// Login -> follower users
	Followers map[string][]string

	// Login -> following users
	Followings map[string][]string

	// Members, Login -> user info
	Members map[string]*github.User
}

const RepoDataFileName string = "gitdog_data.json"

func NewRepo(owner string, repository string) *GitHubRepo {
	return &GitHubRepo{
		Owner:      owner,
		Repository: repository,
	}
}

func (r *GitHubRepo) SetContributors(cs []*github.Contributor) {
	r.Contributors = make(map[string]*github.Contributor)
	for _, c := range cs {
		r.Contributors[*c.Login] = c
	}

	r.Followers = make(map[string][]string)
	r.Followings = make(map[string][]string)

	for _, c := range r.Contributors {
		r.Followers[*c.Login] = make([]string, 0)
		r.Followings[*c.Login] = make([]string, 0)
	}
}

func (r *GitHubRepo) AppendFollowers(c *github.Contributor, followers []*github.User) {
	for _, u := range followers {
		r.Followers[*c.Login] = append(r.Followers[*c.Login], *u.Login)
		r.Followings[*u.Login] = append(r.Followings[*u.Login], *c.Login)
	}
}

func (r *GitHubRepo) AppendFollowings(c *github.Contributor, followings []*github.User) {
	for _, u := range followings {
		r.Followings[*c.Login] = append(r.Followings[*c.Login], *u.Login)
		r.Followers[*u.Login] = append(r.Followers[*u.Login], *c.Login)
	}
}

func (r *GitHubRepo) SetMembers(cs []*github.User) {
	r.Members = make(map[string]*github.User)
	for _, c := range cs {
		r.Members[*c.Login] = c
	}
}

/// Serialization

func (r *GitHubRepo) Save(filePath string) {
	file, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		log.Fatal("Failed to convert data to json")
	}

	err = ioutil.WriteFile(filePath, file, 0644)
	if err != nil {
		log.Fatalf("Failed to save data to file %v", filePath)
	}
}

func Load(filePath string) *GitHubRepo {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read data from file: %v", filePath)
	}

	data := GitHubRepo{}
	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Fatal("Failed to convert json to GitHubRepo")
	}

	return &data
}
