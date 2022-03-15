/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/10-1:13 PM
File: scanner_test.go
*/

package scanner

import (
	"fmt"
	"io"
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

func TestString(t *testing.T) {
	input := `"zjc"`
	quickScan(input)
}

func TestAssign(t *testing.T) {
	//input := `a = 1 + 1`
	input := `a = (1+2)*3+4*5`
	quickScan(input)
}

func TestForLoop(t *testing.T) {
	//input := `a = 1 + 1`
	input := `for (let i = 0; i < 10; i = i + 1) {
	a = a + 1
}`
	quickScan(input)
}

func TestCompare(t *testing.T) {
	inputs := []string{
		`let a = 2++;`,
		``,
		//`<=`,
		//`>`,
		//`<=`,
	}
	for _, input := range inputs {
		quickScan(input)
	}
}

func quickScan(input string) {
	s := Scanner{}
	s.Init(input, func(line, col uint, msg string) {
		log.Println("line:", line, "col:", col, "msg:", msg)
	})
	fmt.Println(input)
	for s.NextTok(); s.err != io.EOF; s.NextTok() {
		fmt.Println("token:", token.Map[s.Type], ", literal:", s.Literal)
	}
	fmt.Println("-----------------")
}
