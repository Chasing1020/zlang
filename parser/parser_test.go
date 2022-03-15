/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/12-12:22 AM
File: parser_test.go
*/

package parser

import (
	"fmt"
	"testing"
	"zlang/scanner"
	"zlang/token"
)

func TestParser(t *testing.T) {
	p := Parser{
		Scanner: scanner.Scanner{},
		errs:    nil,
		curTok:  token.Token{},
		peekTok: token.Token{},
	}
	p.Init("function(a, b) { let zjc = 1+1; return a + b;}")
	file := p.ParseFile()
	fmt.Println(file.Stats)
}

func TestQuickParser(t *testing.T) {
	inputs := []string{
		//"if (true) { let zjc = 1+1; return a + b;}",
		//"(1+2)",
		//"let a = function(a, b) { return a + b; }",
		//"[1, \"a\", 2];",
		//`{a:1, b:2}[a]`,
		//`(1+2)*1+4*2+1*(2*2+1)`,
		//`let a = 5%2`,
		`// This is a comment`,
	}
	for _, input := range inputs {
		quickParser(input)
	}
}

func quickParser(buf string) {
	p := Parser{}
	p.Init(buf)
	file := p.ParseFile()
	fmt.Println("----------------")
	//fmt.Println(file.Stats[0].(*statement.For).Body)
	fmt.Println(file.Stats)
}

func TestForParser(t *testing.T) {
	//buf := `for(let i = 0; i < 5; i = i + 1) {}`
	buf := `let a = function(i) {
				if (i == 0) {
					return 1;
				}
				else {
					return i * function(i-1);
				}
			}`
	//buf := `if(true) {a = a + 2;} `
	quickParser(buf)
}
