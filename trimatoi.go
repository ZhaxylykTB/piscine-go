package piscine

func Check2(str string) bool {
	Pass := false
	for _, v := range str {
		if v >= '0' && v <= '9' {
			Pass = true
			break
		}
	}
	return Pass
}

func SignCheck(s string) bool {
	S := false
	for _, v := range s {
		if v >= '0' && v <= '9' {
			break
		}
		if v == '-' {
			S = true
		}
	}
	return S
}

func TrimAtoi(s string) int {
	x := 0
	if Check2(s) == true {
		for _, v := range s {
			if v >= '0' && v <= '9' {
				c := 0
				for i := '1'; i <= v; i++ {
					c++
				}
				x = x*10 + c
			}
		}
	}
	if SignCheck(s) {
		x *= -1
	}
	return x
}
