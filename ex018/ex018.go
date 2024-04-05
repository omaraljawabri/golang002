package main

import "fmt"

func main() {
	var x, a1, r, n, r1, SomaNumeros float64
	x = 2
	fmt.Scan(&a1)
	fmt.Scan(&r)
	fmt.Scan(&n)
	SomaNumeros = a1
	r1 = r
	for n >= x {
		SomaNumeros = SomaNumeros + a1 + r
		r = r + r1
		x = x + 1
	}
	fmt.Println(SomaNumeros)
}
