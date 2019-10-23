package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for i, v := range args {
		if i != 0 {
			runes := []rune(v)
			for _, r := range runes {
				z01.PrintRune(r)
			}
			z01.PrintRune('\n')
		}
	}
}
