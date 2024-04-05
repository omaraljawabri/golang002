package main

import (
	"fmt"
	"math"
)

func main(){
	var Nota float64
	var Conceito string
	fmt.Scan(&Nota)
	if Nota > 9 {
		Conceito = "A"
	} else if Nota > 7.5 {
		if Nota <= 9 {
			Conceito = "B"	
		}
	} else if Nota >= 6 {
		if Nota <= 7.5 {
			Conceito = "C"	
		}
	} else if Nota < 6 {
		Conceito = "D"
	}
	fmt.Println("NOTA = ", math.Round(Nota*100)/100, "CONCEITO =", Conceito)
}