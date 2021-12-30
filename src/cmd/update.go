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
	client2 "github.com/gitdog7/gitdog/src/client"
	"github.com/gitdog7/gitdog/src/config"
	"github.com/gitdog7/gitdog/src/repo"
	log "github.com/google/logger"
	"github.com/spf13/cobra"
	"os"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update working directory, keep github raw data up-to-date.",
	Long:  `update working directory, keep github raw data up-to-date.`,
	Run: func(cmd *cobra.Command, args []string) {

		// parse flags
		workdir, _ := cmd.Flags().GetString("workdir")

		// check working directory exists
		if _, err := os.Stat(workdir); os.IsNotExist(err) {
			// path exists
			log.Fatalf("working directory: %v doesn't exists.", workdir)
		}

		// load config file
		configFileName := workdir + "/" + config.ConfigFileName
		config := config.Load(configFileName)
		if config == nil {
			log.Fatalf("failed to load config file in %v", configFileName)
		}

		// validate the github client
		client := client2.NewGitHubClient(config.RepoOwner, config.RepoName, config.Token)
		if config.Token == "" {
			log.Warningf("WARNING! github token is empty, the rate limit allows for up to 60 requests per hour.\n")
			log.Warningf("It is recommended that you obtain access token as soon as possible.\n")
			log.Warningf("https://docs.github.com/cn/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token\n")
		}

		if !client.IsClientValid() {
			log.Fatalf("github client is not ready.")
		}

		// begin to update data
		repository := repo.BuildGitHubRepo(client)
		repository.Save(workdir + "/" + repo.RepoDataFileName)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	updateCmd.PersistentFlags().StringP("workdir", "w", ".", "working directory")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
