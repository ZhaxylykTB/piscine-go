package piscine

func StrRev(s string) string {
	runes := []rune(s)
	k := 0
	for i := range s {
		if i == i {
			k++
		}
	}
	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
