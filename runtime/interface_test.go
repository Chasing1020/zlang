/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/12-9:54 PM
File: interface_test.go
*/

package runtime

import (
	"fmt"
	"testing"
)

type A interface {
	a()
}

type B interface {
	A
	b()
}

// AB即实现了A，也实现了B接口
type AB struct{}

func (a *AB) a() {}
func (a *AB) b() {}

func TestName(t *testing.T) {
	var ab1 B = &AB{}
	var ab2 A = &AB{}
	//var aa = ab1.(A)
	//var bb = ab2.(B)
	fmt.Println(ab1 == ab2)
	fmt.Println(equal(ab1, ab2))
}

func equal(a, b A) bool {
	fmt.Println(&a)
	fmt.Println(&b)
	return a == b
}
