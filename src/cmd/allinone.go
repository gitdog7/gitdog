// Package cmd /*
package cmd

import (
	"github.com/gitdog7/gitdog/src/cmd/subcmd"
	"github.com/spf13/cobra"
)

// allinoneCmd represents the allinone command
var allinoneCmd = &cobra.Command{
	Use:   "allinone",
	Short: "generate in-depth report of a given GitHub repository",
	Long:  `generate in-depth report of a given GitHub repository`,
	Run: func(cmd *cobra.Command, args []string) {
		initCmd.Run(cmd, args)
		updateCmd.Run(cmd, args)
		subcmd.GenContributeGraphCmd.Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(allinoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allinoneCmd.PersistentFlags().String("foo", "", "A help for foo")
	allinoneCmd.PersistentFlags().StringP("token", "t", "", "your personal access token.")
	allinoneCmd.PersistentFlags().StringP("repo", "r", "",
		"the GitHub repository you wish to analyze. e.g. go-echarts/go-echarts")
	allinoneCmd.PersistentFlags().StringP("workdir", "w", ".",
		"working directory, data will be stored in this directory.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allinoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
