package main

import "fmt"

func main(){
	var n1, n2, n3 int
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(&n3)
	fmt.Println(n1 + n2 + n3)
	if n1 > 9 {
		fmt.Println("DIGITO INVALIDO")
	} else if n2 > 9 {
		fmt.Println("DIGITO INVALIDO")
	} else if n3 > 9 {
		fmt.Println("DIGITO INVALIDO")
	}

}