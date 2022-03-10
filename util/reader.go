package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"zjclang/scanner"
	"zjclang/token"
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
