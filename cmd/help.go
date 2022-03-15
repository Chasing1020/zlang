/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "help to use",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(`This is my simple language interpreter called zlang.
This language implements some of the features of JavaScript and Python.
It likes a Server-side JavaScript platform (Node.js).

Usage:
	zlang <options> [script.zjc] [arguments]

The commands are:
	run         start run zlang program
	version     show version information
	test        test packages
	fmt         format source files in the current directory
	help        help to use
`)
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
