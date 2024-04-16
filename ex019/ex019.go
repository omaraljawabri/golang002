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
	var n, f int
	var Soma, x float64
	f = 1
	x = 1
	fmt.Scan(&n)
	if n >= 1 {
		for n >= f{
			Soma = Soma + 1/x
			x++
			f++
		}
		fmt.Println(RoundFloat(Soma, 6))
	} else {
		fmt.Println("Número Inválido!")
	}
}
