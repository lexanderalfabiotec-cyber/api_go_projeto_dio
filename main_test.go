package main

import (
	"math"
	"testing"
)

func TestSoma(t *testing.T) {
	if got := Soma(7, 3); got != 10 {
		t.Errorf("Soma(7, 3) = %v; esperado 10", got)
	}
}

func TestSubtrai(t *testing.T) {
	if got := Subtrai(7, 3); got != 4 {
		t.Errorf("Subtrai(7, 3) = %v; esperado 4", got)
	}
}

func TestMultiplica(t *testing.T) {
	if got := Multiplica(7, 3); got != 21 {
		t.Errorf("Multiplica(7, 3) = %v; esperado 21", got)
	}
}

func TestDivide(t *testing.T) {
	got, err := Divide(7, 2)
	if err != nil {
		t.Fatalf("Divide(7, 2) retornou erro inesperado: %v", err)
	}
	if math.Abs(got-3.5) > 1e-9 {
		t.Errorf("Divide(7, 2) = %v; esperado 3.5", got)
	}
}

func TestDividePorZero(t *testing.T) {
	if _, err := Divide(7, 0); err == nil {
		t.Fatal("Divide(7, 0) deveria retornar erro")
	}
}
