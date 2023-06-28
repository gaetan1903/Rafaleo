/*
Copyright © 2023 Gaetan Jonathan BAKARY gaetan.s118@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rafaleo",
	Short: "generate fake data in your database used for load testing",
	Long: `rafaleo is a tool that generates fake data in your database used for load testing.
	`,
	// has an action associated with it:
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		// show config file

	// },
}

var configFile string

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	dir, err := os.Getwd()
	if err != nil {
		dir = "."
	}
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", fmt.Sprintf("config file (default is %s/rafaleofile.yaml)", dir))
}
