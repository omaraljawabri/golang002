package main

import (
	"fmt"
	"math"
)

func main() {
	var SalarioFuncionario, Reajuste float64
	fmt.Scan(&SalarioFuncionario)
	if SalarioFuncionario < 300 {
		Reajuste = SalarioFuncionario * 1.5
	} else {
		Reajuste = SalarioFuncionario * 1.3
	}
	fmt.Println("SALARIO COM REAJUSTE = ", math.Round(Reajuste*100)/100)
}
