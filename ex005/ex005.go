package main

import (
	"fmt"
	"math"
)

func main() {
	var ContaCliente int32
	var ConsumoAgua, ValorConta, Excedente float64
	var TipoConsumidor string
	fmt.Scan(&ContaCliente)
	fmt.Scan(&ConsumoAgua)
	fmt.Scan(&TipoConsumidor)
	switch TipoConsumidor {
	case "R":
		ValorConta = 5 + 0.05*ConsumoAgua
	case "C":
		if ConsumoAgua <= 80 {
			ValorConta = 500
		} else {
			Excedente = ConsumoAgua - 80
			ValorConta = 500 + 0.25*Excedente
		}
	case "I":
		if ConsumoAgua <= 100 {
			ValorConta = 800
		} else {
			Excedente = ConsumoAgua - 100
			ValorConta = 800 + 0.04*Excedente
		}
	}

	fmt.Println("CONTA = ", ContaCliente)
	fmt.Println("VALOR DA CONTA = ", math.Round(ValorConta*100)/100)
}
