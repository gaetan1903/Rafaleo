package core

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Generate(configFile string) {
	// Check if the file exists
	if _, err := os.Stat(configFile); err != nil {
		if os.IsNotExist(err) {
			// print std err
			// add red color

			fmt.Fprintln(os.Stderr, "\033[31mFile", configFile, "does not exist.\033[0m")
			return
		}
		fmt.Fprintln(os.Stderr, "\033[31mError:", err.Error(), "\033[0m")
		return
	}

	// Open the file
	file, err := os.Open(configFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "\033[31mError:", err.Error(), "\033[0m")
		return
	}
	defer file.Close()

	// Read the file contents
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "\033[31mError:", err.Error(), "\033[0m")
		return
	}

	// Print the file contents
	fmt.Println("File contents:", string(contents))
}
