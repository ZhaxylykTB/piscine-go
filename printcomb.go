package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	a := '0'
	b := '0'
	c := '0'

	for a <= '9' {
		for b <= '9' {
			for c <= '9' {
				if a < b && b < c {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(c)
					if a != '7' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				c++
			}
			c = '0'
			b++
		}
		b = '0'
		a++
	}
	z01.PrintRune('\n')
}
