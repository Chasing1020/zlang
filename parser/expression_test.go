/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/12-10:28 AM
File: expression_test.go
*/

package parser

import "testing"

func TestIdentifier(t *testing.T) {
	inputs := []string{
		"a",
		"_",
		"_a",
		"a_",
		"a1",
		"1a",
	}
	for _, input := range inputs {
		quickParser(input)
	}
}

func TestInteger(t *testing.T) {
	inputs := []string{
		"1",
		"123456",
	}
	for _, input := range inputs {
		quickParser(input)
	}
}

func TestString(t *testing.T) {
	inputs := []string{
		"\"\"",
		`"chasing" + "1020"`,
	}
	for _, input := range inputs {
		quickParser(input)
	}
}


func TestBoolean(t *testing.T) {
	inputs := []string {
		"true",
		"falae",
		"2 > 1",
		// TODO: support "true && false",
	}
	for _, input := range inputs {
		quickParser(input)
	}
}