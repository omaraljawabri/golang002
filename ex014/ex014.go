package main

import (
	"fmt"
	"math"
)

func main() {
	var Altura, Aresta, Volume, Ab float64
	fmt.Scan(&Altura)
	fmt.Scan(&Aresta)
	Ab = (3 * math.Pow(Aresta, 2) * math.Sqrt(3))/2
	Volume = Ab*Altura/3
	fmt.Println("O VOLUME DA PIRAMIDE E = ", math.Round(Volume*100)/100, " METROS CUBICOS")
}
