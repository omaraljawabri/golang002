package main

import (
	"fmt"
	"math"
)

func main(){
	var n, QtdDeValores, ValorFahrenheit, ValorCelsius float64
	n = 1
	fmt.Scan(&QtdDeValores)
	for QtdDeValores >= n{
		fmt.Scan(&ValorFahrenheit)
		ValorCelsius = 5*(ValorFahrenheit-32)/9
		fmt.Println(ValorFahrenheit, "FAHRENHEIT EQUIVALE A ", math.Round(ValorCelsius*100)/100, "CELSIUS")
		n++
	}
}