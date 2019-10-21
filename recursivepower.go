package piscine

func RecursivePower(nb, power int) int {
	res := nb
	if power == 0 {
		return 1
	} else if power == 1 {
		return nb
	} else if power < 0 || nb == 0 {
		return 0
	} else {
		return res * RecursivePower(res, power-1)
	}
}
