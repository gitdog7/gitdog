package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type GlobalConfig struct {
	// Token github access token
	Token string

	// RepoOwner
	RepoOwner string

	// RepoName
	RepoName string
}

const ConfigFileName string = "gitdog.json"

/// Serialization

func (r *GlobalConfig) Save(filePath string) {
	file, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		log.Fatal("Failed to convert config to json")
	}

	err = ioutil.WriteFile(filePath, file, 0644)
	if err != nil {
		log.Fatalf("Failed to save config to file %v", filePath)
	}
}

func Load(filePath string) *GlobalConfig {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &GlobalConfig{}
	}

	data := GlobalConfig{}
	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Fatal("Failed to convert json to GitHubRepo")
	}

	return &data
}
