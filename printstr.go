package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	for _, item := range str {
		z01.PrintRune(rune(item))
	}
}
