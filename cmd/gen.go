// Package cmd /*
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
		// load repository
		repository := loadRepository(cmd)

		// visualize
		viz := contributor.ContributeGraphViz{
			Repo: repository,
			TopK: 100,
			Type: "circular",
		}

		// generate contributor graph
		graph := viz.GenerateGraph()
		workdir, _ := cmd.Flags().GetString("workdir")
		outputPath := workdir + "/" +
			strings.Replace(strings.Replace(graph.Title.Title, " ", "_", -1),
				"/", "|", -1) + ".html"
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

// corba cmd init
func init() {
	genCmd.AddCommand(genContributeGraphCmd)
	rootCmd.AddCommand(genCmd)

	genCmd.PersistentFlags().StringP("workdir", "w", "", "working directory")
}

// loadRepository
func loadRepository(cmd *cobra.Command) *repo.GitHubRepo {
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
	return repository
}