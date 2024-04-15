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
	fmt.Printf("O VOLUME DA PIRÂMIDE É = %.2f METROS CÚBICOS", Volume)
}
