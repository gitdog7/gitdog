// Package cmd /*
package cmd

import (
	"github.com/gitdog7/gitdog/src/cmd/subcmd"
	log "github.com/google/logger"
	"github.com/spf13/cobra"
)

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
	genCmd.AddCommand(subcmd.GenContributeGraphCmd)
	rootCmd.AddCommand(genCmd)

	genCmd.PersistentFlags().StringP("workdir", "w", "", "working directory")
}
