package main

import (
	"fmt"
	"math"
)

func main(){
	var x, N float64
	fmt.Scan(&N)
	x = 2
	if N > 5 {
		if N < 2000 {
			if math.Mod(N, 2) == 0 {
				for N >= x{
					fmt.Println(x,"^2 = ", x*x)
					x = x + 2
				}
			}else{
				for (N-1) >= x{
					fmt.Println(x,"^2 = ", x*x)
					x = x + 2					
				}
			}
		}else{
			fmt.Println("Número Inválio!")
		}
	}else{
		fmt.Println("Número Inválido!")
	}
}