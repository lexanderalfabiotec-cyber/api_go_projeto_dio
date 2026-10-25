package main

import (
	"fmt"

	"github.com/lexanderalfabiotec-cyber/go_projeto_dio/pingpong"
)

func main() {
	for _, palavra := range pingpong.GerarSequencia(5) {
		fmt.Println(palavra)
	}
}
