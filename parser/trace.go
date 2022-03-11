/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/11-1:34 PM
File: trace.go
*/

package parser

import "fmt"

// usage: defer p.trace(msg)()
func (p *Parser) trace(msg string) func() {
	p.print(msg + " (")
	return func() {
		if x := recover(); x != nil {
			panic(x) // skip print_trace
		}
		p.print(")")
	}
}

// TODO: implement tracing details for parser
func (p *Parser) print(msg string) {
	fmt.Printf("%5d: %s\n", p.Type, msg)
}
