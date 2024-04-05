package main

import "fmt"

func main() {
	var n, numerodeCasos, Total, Popular, Arquibancada, Geral, Cadeira, RendaDoJogo float64
	n = 1
	fmt.Scan(&numerodeCasos)
	for numerodeCasos >= n {
		fmt.Scan(&Total)
		fmt.Scan(&Popular)
		fmt.Scan(&Geral)
		fmt.Scan(&Arquibancada)
		fmt.Scan(&Cadeira)
		RendaDoJogo = Popular + Geral*5 + Arquibancada*10 + Cadeira*20
		fmt.Println("A RENDA DO JOGO Nº", n, "E = ", RendaDoJogo)
		n++
	}
}
