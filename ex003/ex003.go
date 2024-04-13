package main

import (
	"fmt"
	"strconv"
)

func main(){
	var n1, n2, n3 string
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)
	fmt.Scanln(&n3)
	var z string = n1 + n2 + n3
	var a, b, c, i int
	a, err:= strconv.Atoi(n1)
	b, err = strconv.Atoi(n2)
	c, err = strconv.Atoi(n3)
	i, err = strconv.Atoi(z)
	if err != nil{
		panic(err)
	}
	if a < 10 && b < 10 && c < 10{
		fmt.Print(z,", ",i*i)
	}else{
		fmt.Println("Digito Inválido!")
	}
}