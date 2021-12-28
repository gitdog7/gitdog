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
	"github.com/gitdog7/gitdog/pkg/config"
	"github.com/gitdog7/gitdog/pkg/repo"
	"github.com/gitdog7/gitdog/pkg/visualize/contributor"
	"github.com/gitdog7/gitdog/pkg/visualize/painter"
	"github.com/spf13/cobra"
	"os"
	"strings"

	log "github.com/google/logger"
)

// genContributeGraphCmd represents the gen command
var genContributeGraphCmd = &cobra.Command{
	Use:   "cg",
	Short: "generate contribute graph",
	Long:  `generate contribute graph`,
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

		// load data
		repository := repo.Load(workdir + "/" + repo.RepoDataFileName)

		// visualize
		viz := contributor.ContributeGraphViz{
			Repo: repository,
			TopK: 100,
		}

		// generate contributor graph
		graph := viz.GenerateGraph()
		outputPath := workdir + "/" + strings.Replace(graph.Title.Title, " ", "_", -1) + ".html"
		painter.PaintGraph(graph, outputPath)
		log.Infof("contribute graph: %v generated successfully.", outputPath)

		currentOSWD, _ := os.Getwd()
		log.Infoln("file://" + currentOSWD + "/" + outputPath)
	},
}

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate reports",
	Long:  `generate various reports for given github repo.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Warningln("Please input the report name you want to generate.")
	},
}

func init() {
	genCmd.AddCommand(genContributeGraphCmd)
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	genCmd.PersistentFlags().StringP("workdir", "w", "", "working directory")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
