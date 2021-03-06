/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"zlang/parser"
	"zlang/runtime"

	"github.com/spf13/cobra"
)

// fmtCmd represents the fmt command
var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("zlang run error: no zjc files listed")
			return
		}
		buf, err := os.ReadFile("./" + args[0])
		if err != nil {
			panic(err)
		}
		err = runtime.Run(string(buf))
		if err != nil {
			panic(err)
		}
		p := parser.Parser{}
		p.Init(string(buf))
		file := p.ParseFile()
		fmt.Println(file.Stats)
	},
}

func init() {
	rootCmd.AddCommand(fmtCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fmtCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fmtCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
