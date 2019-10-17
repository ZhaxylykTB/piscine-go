package piscine

func BasicAtoi(s string) int {
	dec := 0
	for _, j := range s {
		ed := 0
		for i := '1'; i <= j; i++ {
			ed = ed + 1
		}
		dec = dec*10 + ed
	}
	return dec
}
