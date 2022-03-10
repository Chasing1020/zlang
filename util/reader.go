/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/10-5:10 PM
File: reader.go
*/

package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"zlang/scanner"
	"zlang/token"
)

func Start() {
	ioScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")
		line := ioScanner.Scan()
		if !line {
			return
		}

		s := scanner.Scanner{}
		s.Init(ioScanner.Text(), func(line, col uint, msg string) {
			log.Println("line:", line, "col:", col, "msg:", msg)
		})
		for s.Next(); s.Tok != token.EOF; s.Next() {
			fmt.Println("token:", token.TokenMap[s.Tok], ", literal:", s.Literal)
		}
	}
}
