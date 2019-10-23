package main

import (
	"github.com/01-edu/z01"
	"os"
)

func Length(str []rune) int {
	c := 0
	for i := range str {
		if i == i {
			c++
		}
	}
	return c
}

func Sol(digit rune) int {
	x := 0
	for i := '1'; i <= digit; i++ {
		if i == i {
			x++
		}
	}
	return x
}

func Sol2(d1, d2 rune) int {
	x := 0
	y := 0
	for i := '1'; i <= d1; i++ {
		if i == i {
			x++
		}
	}
	for i := '1'; i <= d2; i++ {
		if i == i {
			y++
		}
	}
	return x*10 + y
}

func Solut(Arune []rune, len int, i int) []int {
	var arr = []int{0, 0}
	c := 1
	if Arune[i+1] != 32 {
		c++
	}
	if c == 2 {
		if Arune[i] >= '1' && Arune[i] <= '9' && Arune[i+1] >= '0' && Arune[i+1] < '9' {
			arr[1] = i + 2
			arr[0] = Sol2(Arune[i], Arune[i+1])
			return arr
		}

	} else {
		if Arune[i] >= '1' && Arune[i] <= '9' {
			arr[1] = i + 1
			arr[0] = Sol(Arune[i])
			return arr
		}
	}
	return arr
}

func main() {
	args := os.Args
	s := ""
	for i, v := range args {
		if args[i] == "--upper" {
			s = " " + s
			continue
		}
		if i > 0 {
			s += v + " "
		}
	}
	Arune := []rune(s)
	len := Length(Arune)
	for i := 0; i < len-1; i++ {
		if args[1] == "--upper" {
			if Arune[i] == ' ' {
				z01.PrintRune(' ')
				continue
			}
			k := Solut(Arune, len, i)
			if k[0] > 26 {
				z01.PrintRune(32)
			} else {
				z01.PrintRune(rune(97 + k[0] - 33))
			}
			i = k[1]
		} else if Arune[i] >= '0' && Arune[i] <= '9' {
			k := Solut(Arune, len, i)
			if k[0] > 26 {
				z01.PrintRune(32)
			} else {
				z01.PrintRune(rune(97 + k[0] - 1))
			}
			i = k[1]
		} else {
			z01.PrintRune(32)
		}
	}
	z01.PrintRune('\n')
}
