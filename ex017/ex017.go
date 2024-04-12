package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, y float64
	n = 0
	fmt.Scan(&x)
	fmt.Scan(&y)
	if math.Mod(x, 2) == 0 {
		for y > n {
			fmt.Println(x)
			x = x + 2
			n = n + 1
		}
	} else{
		fmt.Println("O Primeiro Numero não é par!")
	}
}
