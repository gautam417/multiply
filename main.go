package main

import (
	"fmt"
)

func solve(width, height, length, mass float64) string {
	volume := width * height * length

	bulky := volume >= 1000000 || width >= 150 || height >= 150 || length >= 150
	heavy := mass >= 20

	if bulky && heavy {
		return "REJECTED"
	} else if bulky || heavy {
		return "SPECIAL"
	} else {
		return "STANDARD"
	}
}

func main() {
	fmt.Println(solve(160, 70, 70, 10))   // SPECIAL
	fmt.Println(solve(120, 100, 100, 10)) // SPECIAL
	fmt.Println(solve(90, 90, 118, 10))   // STANDARD
	fmt.Println(solve(120, 100, 50, 30))  // SPECIAL
	fmt.Println(solve(80, 110, 80, 70))   // SPECIAL
	fmt.Println(solve(70, 80, 90, 10))    // STANDARD
	fmt.Println(solve(100, 120, 60, 19))  // STANDARD
	fmt.Println(solve(150, 70, 70, 10))   // SPECIAL
	fmt.Println(solve(120, 100, 100, 10)) // SPECIAL
	fmt.Println(solve(90, 90, 118, 10))   // STANDARD
	fmt.Println(solve(120, 100, 110, 20)) // REJECTED
	fmt.Println(solve(80, 110, 80, 70))   // SPECIAL
	fmt.Println(solve(70, 80, 90, 10))    // STANDARD
	fmt.Println(solve(100, 150, 60, 30))  // REJECTED
}
