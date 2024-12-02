/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"advent-of-code-2024/internal/day2"
	"fmt"

	"github.com/spf13/cobra"
)

// day2Cmd represents the day2 command
var (
	useDampener *bool
	day2Cmd     = &cobra.Command{
		Use:   "day2",
		Short: "AoC Day 2",
		Long:  `Advent of Code Day 2`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("day2 called with file %s and dampener usage %t\n", args[0], *useDampener)
			day2.Solve(args[0], *useDampener)
		},
	}
)

func init() {
	rootCmd.AddCommand(day2Cmd)
	useDampener = day2Cmd.PersistentFlags().BoolP("useDampener", "d", false, "To use the dampener function")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
