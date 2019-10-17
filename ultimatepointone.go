package piscine

func UltimatePointOne(n ***int) {
	a := 1
	var p1 *int = &a
	var p2 **int = &p1
	var p3 ***int = &p2
	***n = ***p3
}
