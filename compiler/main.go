package main

import (
	"compiler/repl"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
