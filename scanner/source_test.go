/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/10-1:21 PM
File: source_test.go
*/

package scanner

import (
	"fmt"
	"io"
	"log"
	"testing"
)

func TestNextCh(t *testing.T) {
	s := source{}
	s.init("let a = 1+1", func(line, col uint, msg string) {
		log.Println("line:", line, "col:", col, "msg:", msg)
	})
	for s.nextCh(); s.err != io.EOF; s.nextCh() {
		//fmt.Print(string(s.ch))
		fmt.Println(s.pos())
	}
}
