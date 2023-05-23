package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	x := [32]int{}
	if n >= 0 && n <= 9 {
		z01.PrintRune(rune(n + '0'))
	} else {
		for {
			if n == 0 {
				break
			}
			x[n%10]++
			n = n / 10
		}
		for index, z := range x {
			if z != 0 {
				for i := 1; i <= z; i++ {
					z01.PrintRune(rune(index + '0'))
				}
			}
		}
	}
}
