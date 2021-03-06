/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/10-8:36 PM
File: root.go
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"zlang/util"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zlang <options> [script.zjc] [arguments]",
	Short: "This is my tiny language called zlang.\n It is like Server-side JavaScript platform (Node.js).",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`Welcome to zLang v0.0.1.`)
		fmt.Println(`Type "help()" for more information.`)
		util.StartEvaluator()
		fmt.Println("KeyboardInterrupt: EOF, program exited.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SuggestionsMinimumDistance = 1
	rootCmd.SetHelpTemplate(`
This is my simple language interpreter called zlang.
It supports lots of features in JavaScript, which
likes a Server-side JavaScript platform (Node.js).

Usage:
	zlang <options> [script.zjc] [arguments]

The commands are:
	run         start run zlang program
	version     show version information
	test        test packages
	fmt         format source files in the current directory
`)
}
