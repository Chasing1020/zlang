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
	inputs := []string{
		"true",
		"falae",
		"2 > 1",
		// TODO: support "true && false",
	}
	for _, input := range inputs {
		quickParser(input)
	}
}

func TestAssignment(t *testing.T) {
	inputs := []string{
		"a = a + 1",
	}
	for _, input := range inputs {
		quickParser(input)
	}
}

func TestFor(t *testing.T) {
	inputs := []string{
		//"for(let i=0; i <3; let i = i +1 ){i;} ",
		"let a= 1+2*3/4+(1+2)/2",
	}
	for _, input := range inputs {
		quickParser(input)
	}
}

func TestArray(t *testing.T) {
	inputs := []string{
		`let nums = [1, 2, ture, false, "zjc", function(a,b){return a + b;}];
		nums[0]+nums[1];
		nums[2];
		nums[5](1,2);
		`,
	}
	for _, input := range inputs {
		quickParser(input)
	}
}

func TestPalindrome(t *testing.T) {
	buf := `let isPalindrome = function (x) {
    if (x < 0) {
        return false;
    }
    let div = 1;
    for (; x / div >= 10;) {
        div = div * 10;
    }
    for (;x > 0;) {
        let left = x / div;
        let right = x % 10;
        if (left != right) {
            return false;
        }
        x = (x % div) / 10;
        div = div / 100;
    }
    return true;
}

print(isPalindrome(12321));`
	quickParser(buf)
}
