package piscine

func StrLen(str string) int {
	k := 0
	for i := range str {
		if i == i {
			k++
		}
	}
	return k
}
