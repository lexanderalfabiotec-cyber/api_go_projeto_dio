package desafio1

import (
	"reflect"
	"testing"
)

func TestObterDivisiveisPorTres(t *testing.T) {
	tests := []struct {
		name       string
		inicio, fim int
		want       []int
	}{
		{
			name:  "intervalo de 1 a 10",
			inicio: 1,
			fim:    10,
			want:   []int{3, 6, 9},
		},
		{
			name:  "limites divisíveis",
			inicio: 3,
			fim:    9,
			want:   []int{3, 6, 9},
		},
		{
			name:  "intervalo invertido",
			inicio: 10,
			fim:    1,
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ObterDivisiveisPorTres(tt.inicio, tt.fim); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ObterDivisiveisPorTres(%d, %d) = %v; want %v", tt.inicio, tt.fim, got, tt.want)
			}
		})
	}
}
