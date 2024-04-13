package main

import (
	"fmt"
	"math"
)

func main(){
	var n int
	fmt.Scanln(&n)
	temp := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&temp[i])
	}
	for i := 0; i < n; i++ {
		var celsius float64
		celsius = 5*(temp[i]-32)/9
		fmt.Println(temp[i], "FAHRENHEIT EQUIVALE A", math.Round(celsius*100)/100, "CELSIUS" )
	}
}