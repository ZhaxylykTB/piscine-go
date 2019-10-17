package piscine

func Atoi(s string) int {
	dec := 0
	signNumb := 0
	index := 0
	var signV rune
	for i, j := range s {
		ed := 0
		if j == '-' || j == '+' {
			index = i
			signNumb++
			if j == '-' {
				signV = '-'
			}
			if index > 0 {
				return 0
			}
		} else if j < '0' || j > '9' || signNumb > 1 {
			return 0
		}
		for i := '1'; i <= j; i++ {
			ed = ed + 1
		}
		dec = dec*10 + ed
	}
	if signV == '-' {
		return dec * -1
	} else {
		return dec
	}
}
