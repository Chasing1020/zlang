/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/10-5:10 PM
File: reader.go
*/

package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"zlang/object"
	"zlang/parser"
	"zlang/runtime"
	"zlang/scanner"
	"zlang/token"
)

// StartScanner
// TODO: fix keyboard arrows problem: ^[[A ^[[B ^[[C ^[[D
func StartScanner() {
	ioScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")
		line := ioScanner.Scan()
		if !line {
			return
		}

		s := scanner.Scanner{}
		s.Init(ioScanner.Text(), func(line, col uint, msg string) {
			log.Println("line:", line, "col:", col, "msg:", msg)
		})
		for s.NextTok(); s.Type != token.EOF; s.NextTok() {
			fmt.Println("token:", token.Map[s.Type], ", literal:", s.Literal)
		}
	}
}

func StartEvaluator() {
	ioScanner := bufio.NewScanner(os.Stdin)
	env := object.NewEnv()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		<-c
		fmt.Println("\n(To exit, press Ctrl+C again or Ctrl+D)")
		fmt.Print("> ")
		<-c
		os.Exit(0)
	}()

	for {
		fmt.Printf("> ")
		line := ioScanner.Scan()
		if !line {
			return
		}

		p := parser.Parser{}
		p.Init(ioScanner.Text())

		file := p.ParseFile()

		evaluated := runtime.Eval(file, env)
		if evaluated != nil {
			fmt.Println(evaluated.String())
		}
	}
}
