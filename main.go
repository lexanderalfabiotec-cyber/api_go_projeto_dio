package main

import (
	"fmt"

	"github.com/lexanderalfabiotec-cyber/go_projeto_dio/desafio1"
	"github.com/lexanderalfabiotec-cyber/go_projeto_dio/desafio2"
)

func main() {
	fmt.Println("Desafio 1 — números divisíveis por 3:")
	for _, numero := range desafio1.ObterDivisiveisPorTres(1, 100) {
		fmt.Println(numero)
	}

	fmt.Println("\nDesafio 2 — PinPan:")
	for numero := 1; numero <= 100; numero++ {
		fmt.Println(desafio2.ConverterNumero(numero))
	}
}
