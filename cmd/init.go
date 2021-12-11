/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	log "github.com/google/logger"
	client2 "github.com/live77/gitdog/pkg/client"
	config2 "github.com/live77/gitdog/pkg/config"
	"github.com/live77/gitdog/pkg/util"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new working directory.",
	Long:  `Create a new working directory. GitHub repository path and your personal access token is required.`,
	Run: func(cmd *cobra.Command, args []string) {

		// parse flags
		token, _ := cmd.Flags().GetString("token")
		path, _ := cmd.Flags().GetString("repo")
		workdir, _ := cmd.Flags().GetString("workdir")

		// get owner/repo
		res := strings.Split(path, "/")
		if len(res) != 2 {
			log.Fatalf("repo path format is owner/reponame, your input is not valid: %v", path)
		}
		owner := res[0]
		repo := res[1]

		// validate the github client
		client := client2.NewGitHubClient(owner, repo, token)
		if !client.IsClientValid() {
			log.Fatalf("github client is not ready, token %v", token)
		}
		if token == "" {
			log.Warningf("WARNING! github token is empty, the rate limit allows for up to 60 requests per hour.\n")
			log.Warningf("It is recommended that you obtain access token as soon as possible.\n")
			log.Warningf("https://docs.github.com/cn/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token\n")
		}

		// check working directory not exists
		if _, err := os.Stat(workdir); !os.IsNotExist(err) {
			// path exists
			log.Fatalf("working directory: %v already exists.", workdir)
		}

		// create working directory
		if err := os.MkdirAll(workdir, 0755); err != nil {
			log.Fatalf("failed to create working directory: %v\n, %v", workdir, err)
		}

		// download echarts.min.js to work directory
		if err := util.DownloadFile(workdir+"/"+"echarts.min.js",
			"https://live77.github.io/resource/js/echarts.min.js"); err != nil {
			log.Fatalf("failed to download echarts.min.js.")
		}

		// create config file and save it
		config := config2.GlobalConfig{
			Token:     token,
			RepoOwner: owner,
			RepoName:  repo,
		}
		config.Save(workdir + "/gitdog.json")
		log.Infof("working directory created successfully, %v, %v\n", workdir, config)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	initCmd.PersistentFlags().StringP("token", "t", "", "your personal access token.")
	initCmd.PersistentFlags().StringP("repo", "r", "",
		"the GitHub repository you wish to analyze. e.g. go-echarts/go-echarts")
	initCmd.PersistentFlags().StringP("workdir", "w", ".",
		"working directory, data will be stored in this directory.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
