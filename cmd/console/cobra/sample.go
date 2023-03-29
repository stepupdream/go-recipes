package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sampleCmd represents the command1 command
var sampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "A brief description of your command",
	Long:  `A long multi line description of the command`,

	// Contents of processing during command execution
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Flag("toggle1").Value)
		fmt.Println(cmd.Flag("toggle2").Value)
		fmt.Println(cmd.Flag("toggle3").Value)
	},
}

func init() {
	// Spell for adding subcommands
	rootCmd.AddCommand(sampleCmd)

	// Specify command arguments and options(required)
	sampleCmd.Flags().StringP("toggle1", "a", "aaa", "Help message for toggle")
	_ = sampleCmd.MarkFlagRequired("toggle1")

	// Specify command arguments and options
	sampleCmd.Flags().StringP("toggle2", "t", "aaa", "Help message for toggle")
	sampleCmd.Flags().BoolP("toggle3", "s", false, "Help message for toggle")
}
