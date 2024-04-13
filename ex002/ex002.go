package main

import (
	"fmt"
	"math"
)

func main() {
	var n, numerodeCasos, Total, PPopular, PCadeira, PGeral, PArquibancada, RendaDoJogo, Popular, Geral, Arquibancada, Cadeira float64
	n = 1
	fmt.Scan(&numerodeCasos)
	for numerodeCasos >= n {
		fmt.Scan(&Total)
		fmt.Scan(&PPopular)
		fmt.Scan(&PGeral)
		fmt.Scan(&PArquibancada)
		fmt.Scan(&PCadeira)
		Popular = (PPopular/100)*Total
		Geral = (PGeral/100)*Total
		Cadeira = (PCadeira/100)*Total
		Arquibancada = (PArquibancada/100)*Total
		RendaDoJogo = Popular + Geral*5 + Arquibancada*10 + Cadeira*20
		fmt.Println("A RENDA DO JOGO Nº", n, "E = ", math.Round(RendaDoJogo*1000)/1000)
		n++
	}
}
