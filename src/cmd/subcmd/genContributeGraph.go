package subcmd

import (
	"fmt"
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
		repo := loadRepository(cmd)
		workdir, _ := cmd.Flags().GetString("workdir")

		// visualize
		option := contributor.ContributeGraphOption{
			TopK: 100,
			Type: "circular",
		}

		// generate contributor graph (circular)
		graph1 := contributor.GenerateGraph(repo, option)
		outputPath1 := workdir + "/" + getGraphSaveName(graph1.Title.Title) + ".html"
		painter.PaintGraph(graph1, outputPath1)
		log.Infof("contribute graph: %v generated successfully.", outputPath1)

		// generate contribute graph (force)
		option.Type = "force"
		graph2 := contributor.GenerateGraph(repo, option)
		outputPath2 := workdir + "/" + getGraphSaveName(graph2.Title.Title) + ".html"
		painter.PaintGraph(graph2, outputPath2)
		log.Infof("contribute graph: %v generated successfully.", outputPath2)

		currentOSWD, _ := os.Getwd()
		fmt.Println("-----------------------------")
		fmt.Println("file://" + currentOSWD + "/" + outputPath1)
		fmt.Println("file://" + currentOSWD + "/" + outputPath2)
	},
}
