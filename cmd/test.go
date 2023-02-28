/*
Copyright © 2023 Gaetan Jonathan BAKARY gaetan.s118@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your test",
	Long:  `a long description test`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("it works!")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
