package subcmd

import (
	"github.com/gitdog7/gitdog/src/visualize/contributor"
	"github.com/gitdog7/gitdog/src/visualize/painter"
	log "github.com/google/logger"
	"github.com/spf13/cobra"
	"os"
)

// GenContributeGraphCmd generate contributes graphs
var GenContributeGraphCmd = &cobra.Command{
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
		outputPath := workdir + "/" + getGraphSaveName(graph.Title.Title) + ".html"
		painter.PaintGraph(graph, outputPath)
		log.Infof("contribute graph: %v generated successfully.", outputPath)

		currentOSWD, _ := os.Getwd()
		log.Infoln("file://" + currentOSWD + "/" + outputPath)
	},
}
