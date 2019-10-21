package piscine

func IterativePower(nb int, power int) int {

	rp := nb

	if power < 0 || power > 50 {
		return 0
	}
	if power == 0 {
		return 1
	}
	for i := 1; i < power; i++ {
		rp = rp * nb
	}
	return rp
}
