package cmd

import (
	"github.com/gaetan1903/rafaleo/core"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate data in your database",
	Long:  "run rafaleo to generate data in your database",
	Run: func(cmd *cobra.Command, args []string) {
		core.Generate(configFile)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
