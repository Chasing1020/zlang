/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/10-8:40 PM
File: run.go
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"zlang/runtime"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Compile and run a source file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("zlang run error: no go files listed")
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
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
