package main

import (
	"flag"
	"github.com/mishaprokop4ik/brainfuck/internal/input"
	"github.com/mishaprokop4ik/brainfuck/internal/interpreter"
	"log"
)

var programPath = flag.String("f", "", "brainfuck program file path")

const defaultDataIn string = ">++++++++[<+++++++++>-]<.>++++[<+++++++>-]<+.+++++++..+++.>>++++++[<+++++++>-]<++.------------.>++++++[<+++++++++>-]<+.<.+++.------.--------.>>>++++[<++++++++>-]<+.\n"

func main() {
	flag.Parse()
	fromFile := true
	if *programPath == "" {
		*programPath = defaultDataIn
		fromFile = false
	}
	dataIn, err := input.NewFile(*programPath, fromFile)
	if err != nil {
		log.Fatal(err)
	}

	interpreter.Interpret(dataIn.OutText())
)
