// Package desafio2 contém a conversão numérica da brincadeira PinPan.
package desafio2

import "strconv"

// ConverterNumero retorna "PinPan" para múltiplos de 15, "Pin" para os
// demais múltiplos de 3, "Pan" para os demais múltiplos de 5 e o próprio
// número convertido em texto nos outros casos.
func ConverterNumero(n int) string {
	switch {
	case n%15 == 0:
		return "PinPan"
	case n%3 == 0:
		return "Pin"
	case n%5 == 0:
		return "Pan"
	default:
		return strconv.Itoa(n)
	}
}
