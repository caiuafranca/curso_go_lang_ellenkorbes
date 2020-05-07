package main

import (
	"fmt"
)

//criando Tipo de dado personalidado
type cursoEllen int

var (
	x int
	//Y exemplo de variavel publica
	Y string
	//Z exemplo de variavel Publica
	Z bool
	//C exemplo de criação de Tipo
	C cursoEllen
)

//Desafio Ninja iniciante
func main() {
	x = 42
	Y = "James Bond"
	Z = true
	C = 20
	var x2 = cursoEllen(x)

	fmt.Printf("Bem Vindo ao Mundo dos gophers, %d, %s, %v \n", x, Y, Z)
	saida := fmt.Sprintf("%T", C)
	fmt.Printf("o dado e do tipo %s e tem valor %d \n", saida, C)

	fmt.Println("Cast do tipo")
	saida2 := fmt.Sprintf("%T", x2)
	fmt.Printf("o dado e do tipo %s e tem valor %d \n", saida2, x2)

	fmt.Println("Vamos ao Nível Ninja - Level 01")

}
