package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	runes := []rune(args[0])
	for _, v := range runes {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
