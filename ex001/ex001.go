package main

import (
	"fmt"
	"math"
)

func main() {
	var nota1, nota2, nota3, MEDIA float64
	fmt.Scanln(&nota1, &nota2, &nota3)
	MEDIA = (nota1 + nota2 + nota3) / 3
	if MEDIA >= 6 {
		fmt.Println("MEDIA = ", math.Round(MEDIA*100)/100)
		fmt.Println("APROVADO")
	} else if MEDIA <= 6 {
		fmt.Println("MEDIA = ", math.Round(MEDIA*100)/100)
		fmt.Println("REPROVADO")
	}
}
