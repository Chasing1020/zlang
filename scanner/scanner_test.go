/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/10-1:13 PM
File: scanner_test.go
*/

package scanner

import (
	"fmt"
	"log"
	"testing"
	"zlang/token"
)

func TestScannerNext(t *testing.T) {
	tests := []string{
		"123+1*(23+34)",
		"   let  a = 1 +11 23   14143245123531451   ",
		"_ast",
		"1a",
	}
	for _, test := range tests {
		quickScan(test)
	}
}

func TestLet(t *testing.T) {
	input := "function(a) { return 1}; "
	quickScan(input)
}

func quickScan(input string) {
	s := Scanner{}
	s.Init(input, func(line, col uint, msg string) {
		log.Println("line:", line, "col:", col, "msg:", msg)
	})
	fmt.Println(input)
	for s.NextTok(); s.Type != token.EOF; s.NextTok() {
		fmt.Println("token:", token.TokenMap[s.Type], ", literal:", s.Literal)
	}
	fmt.Println("-----------------")
}