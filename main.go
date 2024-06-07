package main

import (
	"fmt"

	"github.com/oduortoni/syntax-suggest/lib/core"
)

func main() {
	output := core.CorrectAndRun("hello.txt", "hello.syntax")
	fmt.Printf("%s\n", output)
}
