package main

import (
	"fmt"
)

func main() {
	var n, numerodeCasos int
	var Total, PPopular, PCadeira, PGeral, PArquibancada, RendaDoJogo, Popular, Geral, Arquibancada, Cadeira float64
	n = 1
	fmt.Scanln(&numerodeCasos)
	for numerodeCasos >= n {
		fmt.Scanln(&Total, &PPopular, &PGeral, &PArquibancada, &PCadeira)
		Popular = (PPopular/100)*Total
		Geral = (PGeral/100)*Total
		Cadeira = (PCadeira/100)*Total
		Arquibancada = (PArquibancada/100)*Total
		RendaDoJogo = Popular + Geral*5 + Arquibancada*10 + Cadeira*20
		fmt.Printf("A RENDA DO JOGO Nº %d É = %.2f \n", n, RendaDoJogo)
		n++
	}
}
