package piscine

import "github.com/01-edu/z01"

func PrintNum(num int) {
	i := '0'
	if num == 0 {
		z01.PrintRune('0')
		return
	}
	for j := 1; j <= num%10; j++ {
		i++
	}
	for j := -1; j >= num%10; j-- {
		i++
	}
	if num/10 != 0 {
		PrintNum(num / 10)
	}
	z01.PrintRune(i)
	return
}

func PrintNbr(n int) {

	if n < 0 {
		z01.PrintRune('-')
	}
	PrintNum(n)
	/*if n == 0 {
		z01.PrintRune('0')
	} else {
		var num int = 0
		var minus rune
		var count = 1
		if n < 0 {
			minus = '-'
			n *= (-1)
		}
		var reverseNum int = n % 10
		n /= 10
		z01.PrintRune(IntToRune(reverseNum))
		for n != 0 {
			reverseNum *= 10
			num = n % 10
			z01.PrintRune(IntToRune(num))
			reverseNum += num
			n /= 10
			count++
		}
		z01.PrintRune('\n')
		fmt.Println(count)
		if minus == '-' {
			z01.PrintRune('-')
		}
		var numnum int
		for j := 0; j <= count; j++ {
			numnum = reverseNum % 10
			z01.PrintRune(IntToRune(numnum))
			reverseNum = reverseNum / 10
		}
		//var a int = reverseNum
		//var r int = a % 10
		//z01.PrintRune(IntToRune(r))
		//a /= 10
		//for a > 0 {
		//	r *= 10
		//	num = a % 10
		//	z01.PrintRune(IntToRune(num))
		//	r += num
		//	a /= 10
		//}
		//z01.PrintRune('\n')
		//for reverseNum != 0 {
		//
		//	num = reverseNum % 10
		//	z01.PrintRune(IntToRune(num))
		//	reverseNum = reverseNum / 10
		//}
	}
	*/
}
