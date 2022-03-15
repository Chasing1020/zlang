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
	env := object.NewEnv()
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

func TestBoolean(t *testing.T) {
	inputs := []string{
		`
		let a =  -2 >= -1; 
		a;
		`,
	}
	for _, input := range inputs {
		quickEval(input)
	}
}

func TestAssignment(t *testing.T) {
	inputs := []string{
		//`let a = 1; a = 2; a;`,
		//`let a = 0;
		//for (let i = 0; i <= 100; i = i + 1) {a = a + i;}
		//a;`,
		`let a = function(i) {
		if (i == 0) {return 1;}
		else {return i * a(i-1);}
		}
		a(5);`,
	}
	for _, input := range inputs {
		quickEval(input)
	}
}

func TestArray(t *testing.T) {
	inputs := []string{
		//`let nums = [1, 2,true, "zjc", function(a,b){return a + b;}];
		//nums[0]+nums[1];
		//nums[3](1,2);
		//`,
		//`let a={"name":"zjc"};
		//a["name"] = "chasing";
		//a["name"]`,
		//`let a = eval(1+1*2);
		//println(a)
		//println(string(a))`,
		//`// This is a comment that
		//let a = 10;`,
		`3!=5`,
	}
	for _, input := range inputs {
		quickEval(input)
	}
}

func TestTwoSum(t *testing.T) {
	buf := `let nums = [2, 7, 11, 15];
let target = 9;
let dict = {};
for (let i = 0; i < len(nums); i = i + 1) {
   if (dict[nums[i]] == null) {
       dict[target - nums[i]] = i;
   } else {
       printf("Answer found: [%s] [%s]", i, dict[nums[i]]);
   }
}`
	quickEval(buf)
}

func TestFor(t *testing.T) {
	buf := `let i = 0;	
	for(;i<5;){i=i+1;}	
`
	quickEval(buf)
}

func quickEval(buf string) {
	p := parser.Parser{}
	p.Init(buf)
	file := p.ParseFile()
	env := object.NewEnv()
	eval := Eval(file, env)
	log.Println(eval)
}
