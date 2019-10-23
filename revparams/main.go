package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	counter := 0
	for range args {
		counter++
	}
	for i := counter - 1; i > 0; i-- {
		runes := []rune(args[i])
		for _, r := range runes {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
