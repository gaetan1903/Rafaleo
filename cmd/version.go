/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/gaetan1903/rafaleo/constant"
	"github.com/gaetan1903/rafaleo/utils"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Long:  "show current version of rafaleo",
	Run: func(cmd *cobra.Command, args []string) {
		os.Stdout.WriteString(" ⭐ \033[32m")
		utils.TypePrint(constant.Version["version"])
		os.Stdout.WriteString("\033[0m")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
