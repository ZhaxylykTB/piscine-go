package piscine

func StrLen3(s string) int {
	counter := 0
	for i := range s {
		if i == i {
			counter++
		}
	}
	return counter
}

func NRune(s string, n int) rune {
	Arune := []rune(s)
	len := StrLen3(s)
	if n <= len && n-1 >= 0 {
		return Arune[n-1]
	} else {
		return 0
	}
}
