package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func toint(nb int) {
	x := '0'
	if nb/10 > 0 {
		toint(nb / 10)
	}
	for i := 0; i < nb%10; i++ {
		x++
	}
	z01.PrintRune(x)
}

func tostr(s string) {
	a := []rune(s)
	for index := range a {
		z01.PrintRune(a[index])
	}
}

func main() {
	points := &point{}
	setPoint(points)

	tostr("x = ")
	toint(points.x)
	tostr(", y = ")
	toint(points.y)
	z01.PrintRune('\n')
}
