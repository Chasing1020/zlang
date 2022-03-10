package token

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToken(t *testing.T) {
	tok := Token(1)
	v := reflect.ValueOf(tok)
	ty := reflect.TypeOf(&tok)
	fmt.Println(v, ty)
}

func TestMakeTokenMap(t *testing.T) {
	m := map[uint]string{
		0:  "UNKNOWN",
		1:  "EOF",
		2:  "Ident",
		3:  "Int",
		4:  "String",
		5:  "Operator",
		6:  "Assign",
		7:  "Plus",
		8:  "Minus",
		9:  "Bang",
		10: "Star",
		11: "Slash",
		12: "Eql",
		13: "Neq",
		14: "Lss",
		15: "Leq",
		16: "Gtr",
		17: "Geq",
		18: "Lparen",
		19: "Lbrack",
		20: "Lbrace",
		21: "Rparen",
		22: "Rbrack",
		23: "Rbrace",
		24: "Comma",
		25: "Semi",
		26: "Colon",
		27: "Dot",
		28: "DotDotDot",
		29: "Function",
		30: "True",
		31: "False",
		32: "If",
		33: "Else",
		34: "Let",
		35: "Return",
	}
	for i := 0; i < 36; i++ {
		fmt.Printf("\"%s\",",m[uint(i)])
	}
}
