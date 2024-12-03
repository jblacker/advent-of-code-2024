/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"advent-of-code-2024/internal/day3"
	"fmt"

	"github.com/spf13/cobra"
)

var useConditionals *bool

// day3Cmd represents the day3 command
var day3Cmd = &cobra.Command{
	Use:   "day3",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("day3 called with path: %s & use conditionals: %t\n", args[0], *useConditionals)
		day3.Solve(args[0], *useConditionals)

	},
}

func init() {
	rootCmd.AddCommand(day3Cmd)
	useConditionals = day3Cmd.PersistentFlags().BoolP("conditional", "c", false, "To use conditionals in the input")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
