package piscine

func StrLen2(s string) int {
	counter := 0
	for i := range s {
		if i == i {
			counter++
		}
	}
	return counter
}

func LastRune(s string) rune {
	Arune := []rune(s)
	len := StrLen2(s)
	return Arune[len-1]
}
