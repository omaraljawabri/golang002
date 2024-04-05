package main

import (
	"fmt"
	"math"
)

func main(){
	var ValorA, ValorB, ValorC, ValorD, determinante float64
	fmt.Scan(&ValorA)
	fmt.Scan(&ValorB)
	fmt.Scan(&ValorC)
	fmt.Scan(&ValorD)
	determinante = ValorA*ValorD - ValorB*ValorC
	fmt.Println("O VALOR DO DETERMINANTE E = ", math.Round(determinante*100)/100)
}