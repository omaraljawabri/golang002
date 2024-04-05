package main

import (
	"fmt"
	"math"
)

func main() {
	var Numero float64
	fmt.Scan(&Numero)
	if math.Mod(Numero, 3) == 0 {
		if math.Mod(Numero, 5) == 0 {
			fmt.Println("O NUMERO E DIVISIVEL")
		}
	} else {
		fmt.Println("O NUMERO NAO E DIVISIVEL")
	}
}
