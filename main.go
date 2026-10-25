package main

import "errors"

// Soma retorna a soma de dois números.
func Soma(a, b float64) float64 {
	return a + b
}

// Subtrai retorna a diferença entre dois números.
func Subtrai(a, b float64) float64 {
	return a - b
}

// Multiplica retorna o produto de dois números.
func Multiplica(a, b float64) float64 {
	return a * b
}

// Divide retorna a divisão de a por b ou um erro se b for zero.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero")
	}
	return a / b, nil
}

func main() {}
