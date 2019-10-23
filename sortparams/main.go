package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	var runes []rune
	res := ""
	counter := 0
	for i, v := range args {
		if i != 0 {
			counter++
			res += v
		}
	}
	runes = []rune(res)
	for i := 0; i < counter-1; i++ {
		for j := i + 1; j < counter; j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	for _, v := range runes {
		z01.PrintRune(v)
		z01.PrintRune('\n')
	}
}
