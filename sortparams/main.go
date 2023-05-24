package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	fn := os.Args[1:]
	var len int
	for index := range fn {
		len = index + 1
	}
	for i := 0; i < len(fn)-1; i++ {
		for j := 0; j < (len - 1 - i); j++ {
			if fn[j] > fn[j+1] {
				fn[j], fn[j+1] = fn[j+1], fn[i]
			}
		}
	}
	for _, x := range fn {
		for _, y := range x {
			z01.PrintRune(y)
		}
		z01.PrintRune('\n')
	}
}
