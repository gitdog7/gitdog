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
	"github.com/spf13/cobra"
)

// allinoneCmd represents the allinone command
var allinoneCmd = &cobra.Command{
	Use:   "allinone",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		initCmd.Run(cmd, args)
		updateCmd.Run(cmd, args)
		genContributeGraphCmd.Run(cmd, args)
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
