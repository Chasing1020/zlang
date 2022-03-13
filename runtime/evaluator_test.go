/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/12-8:51 PM
File: evaluator_test.go
*/

package runtime

import (
	"fmt"
	"log"
	"testing"
	"zlang/object"
	"zlang/parser"
)

func TestEvaluator(t *testing.T) {
	buf := `let a = 1;`
	p := parser.Parser{}
	p.Init(buf)
	file := p.ParseFile()
	env := object.NewEnvironment()
	eval := Eval(file, env)
	log.Println(eval)
}

func QuickParser(buf string) {
	p := parser.Parser{}
	p.Init(buf)
	file := p.ParseFile()
	fmt.Println("----------------")
	fmt.Println(file.Stats)
}

func TestAssignment(t *testing.T) {
	inputs := []string{
		//`let a = 1; a = 2; a;`,
		`
		let a = 0; 
		for (let i = 0; i < 10; i = i + 1) {a = a + i;}
		a;
		`,
	}
	for _, input := range inputs {
		quickEval(input)
	}
}

func TestArray(t *testing.T) {
	inputs := []string{
		`let nums = [1, 2, 3, 4, 5];
		nums[0];
		`,
	}
	for _, input := range inputs {
		quickEval(input)
	}
}

func quickEval(buf string) {
	p := parser.Parser{}
	p.Init(buf)
	file := p.ParseFile()
	env := object.NewEnvironment()
	eval := Eval(file, env)
	log.Println(eval)
}

