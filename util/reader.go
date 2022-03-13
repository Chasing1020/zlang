/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/10-5:10 PM
File: reader.go
*/

package util

import (
	"bufio"
	"fmt"
	"github.com/mattn/go-tty"
	"io"
	"log"
	"os"
	"zlang/object"
	"zlang/parser"
	"zlang/runtime"
	"zlang/scanner"
	"zlang/token"
)

func Start() {
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

func StartEvaluator(in io.Reader, out io.Writer) {
	//s := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf("> ")

		r := bufio.NewReader(os.Stdin)
		str, _ := r.ReadString('\n')

		p := parser.Parser{}
		//p.Init(s.Text())
		p.Init(str)
		file := p.ParseFile()

		evaluated := runtime.Eval(file, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.String())
			io.WriteString(out, "\n")
		}
	}
}

func StartTTY() {
	t, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer t.Close()
	env := object.NewEnvironment()
	for {

		str, err := t.ReadString()
		if err != nil {
			log.Fatal(err)
		}

		p := parser.Parser{}
		//p.Init(s.Text())
		p.Init(str)
		file := p.ParseFile()

		evaluated := runtime.Eval(file, env)
		if evaluated != nil {
			io.WriteString(os.Stdout, evaluated.String())
			io.WriteString(os.Stdout, "\n")
		}
		// handle key event
	}
}
