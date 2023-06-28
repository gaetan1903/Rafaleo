package cmd

import (
	"fmt"
	"os"

	"github.com/gaetan1903/rafaleo/constant"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a config file",
	Long:  "create a template yaml for config file",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var _configFile string
		if len(args) > 0 {
			_configFile = args[0]
		} else {
			_configFile = "rafaleofile.yaml"
		}

		file, err := os.Create(_configFile)
		if err != nil {
			fmt.Println("Failed to create file:", err)
			return
		}
		defer file.Close()

		_, err = file.Write([]byte(constant.TemplateContent))
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			return
		}

		fmt.Println("Content written to file successfully!")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
