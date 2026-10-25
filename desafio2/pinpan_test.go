package desafio2

import "testing"

func TestConverterNumero(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{n: 1, want: "1"},
		{n: 3, want: "Pin"},
		{n: 5, want: "Pan"},
		{n: 15, want: "PinPan"},
		{n: 30, want: "PinPan"},
		{n: -3, want: "Pin"},
		{n: 0, want: "PinPan"},
	}

	for _, tt := range tests {
		if got := ConverterNumero(tt.n); got != tt.want {
			t.Errorf("ConverterNumero(%d) = %q; want %q", tt.n, got, tt.want)
		}
	}
}
