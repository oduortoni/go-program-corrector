package main

import (
	"gpcorrect/lib/core"
)

func main() {
	output := core.CorrectAndRun("programs/hello.txt")
	println(output)
}
