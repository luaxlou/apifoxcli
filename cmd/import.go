/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/luaxlou/apifoxcli/lib/apifox"
	"github.com/spf13/cobra"
	"os"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import api docs to apifox",
	Run: func(cmd *cobra.Command, args []string) {

		projectId, _ := cmd.Flags().GetString("project-id")
		docsFilePath, _ := cmd.Flags().GetString("docs-file-path")
		content, _ := os.ReadFile(docsFilePath)

		apifox.ImportOpenApi(projectId, string(content))

	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().StringP("project-id", "", "", "project id")
	importCmd.Flags().StringP("docs-file-path", "f", "", "api docs file path")
	importCmd.Flags().StringP("access-token", "t", "", "apifox access token")
}
