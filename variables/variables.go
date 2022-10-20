package main

import "fmt"

func main() {
	var numero int
	numero = 35
	fmt.Println(numero)
	numero = 40
	fmt.Println(numero)

	nombre := "Phellipe"
	fmt.Println(nombre)

	nombre, numero = "Rodrigues", 32

	fmt.Println(nombre, numero)
}
