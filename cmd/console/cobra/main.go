package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// main represents the base command when called without any subcommands
func main() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "go-recipes",
}
