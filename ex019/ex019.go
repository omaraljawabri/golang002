package main

import (
	"fmt"
	"math"
)

func RoundFloat(val float64, precision uint) float64{
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio)/ratio
}

func main() {
	var x, n, Soma float64
	x = 1
	fmt.Scan(&n)
	if n > 1 {
		for n >= x {
			Soma = Soma + 1/x
			x = x + 1
		}
	} else {
		fmt.Println("Numero Invalido!")
	}
fmt.Println(RoundFloat(Soma, 6))
}
