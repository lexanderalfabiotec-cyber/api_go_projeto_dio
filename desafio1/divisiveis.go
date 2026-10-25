// Package desafio1 contém funções para localizar números divisíveis por 3.
package desafio1

// ObterDivisiveisPorTres retorna os números divisíveis por 3 no intervalo
// inclusivo entre inicio e fim. Se inicio for maior que fim, retorna uma
// fatia vazia.
func ObterDivisiveisPorTres(inicio, fim int) []int {
	var divisiveis []int
	for numero := inicio; numero <= fim; numero++ {
		if numero%3 == 0 {
			divisiveis = append(divisiveis, numero)
		}
	}
	return divisiveis
}
