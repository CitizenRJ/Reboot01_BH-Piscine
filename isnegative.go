package piscine

import "fmt"

func IsNegative(nb int) {
	if nb < 0 {
		fmt.Print("T")
	} else {
		fmt.Print("F")
	}
}

func main() {
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
	fmt.Println()
}
