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
	s := Scanner{}
	s.Init("123+1*(23+34)", func(line, col uint, msg string) {
		log.Println("line:", line, "col:", col, "msg:", msg)
	})
	for s.Next(); s.Tok != token.EOF; s.Next() {
		fmt.Println("token:", token.TokenMap[s.Tok], ", literal:", s.Literal)
	}
}
