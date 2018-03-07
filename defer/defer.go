package main

import (
	"fmt"
)

func obtervalorProvado(numero int) int {
	defer fmt.Println("fim!")

	if numero >= 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero
}

func main() {
	fmt.Println(obtervalorProvado(6000))
	fmt.Println(obtervalorProvado(4000))
}
