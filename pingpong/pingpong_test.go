package pingpong

import (
	"reflect"
	"testing"
)

func TestGerarSequencia(t *testing.T) {
	got := GerarSequencia(3)
	want := []string{"ping", "pong", "ping", "pong", "ping", "pong"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GerarSequencia(3) = %v; want %v", got, want)
	}
}

func TestGerarSequenciaSemPares(t *testing.T) {
	for _, pares := range []int{0, -1} {
		got := GerarSequencia(pares)
		if len(got) != 0 {
			t.Errorf("GerarSequencia(%d) = %v; want sequência vazia", pares, got)
		}
	}
}
