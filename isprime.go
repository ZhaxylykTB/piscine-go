package piscine

func IsPrime(nb int) bool {
	if nb <= 0 || nb > 10000000000 || nb == 1 {
		return false
	} else {
		prime := true
		for j := 2; j*j <= nb; j++ {
			if nb%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			return true
		} else {
			return false
		}
	}
}
