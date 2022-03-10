/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/10-20:40 PM
File: version.go
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show the version of the zlang",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("zlang version 0.0.1")
	},
}


func init() {
	rootCmd.AddCommand(versionCmd)
}



