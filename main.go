/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/10-12:30 PM
File: main.go
*/

package main

import (
	"os"
	"zlang/util"
)

func main() {
	util.StartEvaluator(os.Stdin, os.Stdout)
	//util.StartTTY()
	//cmd.Execute()
}
