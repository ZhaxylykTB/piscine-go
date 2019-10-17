package piscine

func BasicAtoi2(s string) int {
	dec := 0
	for _, j := range s {
		ed := 0
		if j < '0' || j > '9' {
			return 0
		}
		for i := '1'; i <= j; i++ {
			ed = ed + 1
		}
		dec = dec*10 + ed
	}
	return dec
}
