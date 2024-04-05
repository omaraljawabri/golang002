package main

import (
	"fmt"
	"math"
)

func main(){
	var Fahrenheit, QtdChuva, Celsius, Polegadas float64
	fmt.Scan(&Fahrenheit)
	fmt.Scan(&Polegadas)
	Celsius = (5*Fahrenheit - 160)/9
	QtdChuva = Polegadas*25.4
	fmt.Println("O VALOR EM CELSIUS E = ", math.Round(Celsius*100)/100)
	fmt.Println("A QUANTIDADE DE CHUVA E = ", math.Round(QtdChuva*100)/100)
}