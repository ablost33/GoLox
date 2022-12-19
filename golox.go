package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	numArgs := len(args)

	switch {
	case numArgs > 1:
		fmt.Println("Usage: golox [script]")
		os.Exit(64)
	case numArgs == 1:
		err := runFile(args[0])
		if err != nil {
			os.Exit(1)
		}
	default:
		runPrompt()
	}
}

func runFile(path string) error {
	buf, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("[runFile] os.ReadFile failed: %s", err)
		return err
	}
	run(string(buf[:]))
	return nil
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Printf("[runPrompt] Error reading input: %s", err)
		}
		if line == nil || len(line) == 0 {
			return
		}

		run(string(line[:]))
	}
}
