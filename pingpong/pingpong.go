// Package pingpong demonstra alternância entre goroutines usando canais.
package pingpong

// GerarSequencia cria pares alternados de "ping" e "pong".
//
// Uma goroutine é responsável por cada palavra. Os canais pingTurn e pongTurn
// funcionam como um mecanismo de passagem de turno, garantindo a ordem sem
// depender de temporização (sleep).
func GerarSequencia(pares int) []string {
	if pares <= 0 {
		return []string{}
	}

	palavras := make(chan string)
	pingTurn := make(chan struct{})
	pongTurn := make(chan struct{})

	go func() {
		for i := 0; i < pares; i++ {
			<-pingTurn
			palavras <- "ping"
			pongTurn <- struct{}{}
		}
	}()

	go func() {
		for i := 0; i < pares; i++ {
			<-pongTurn
			palavras <- "pong"
			if i < pares-1 {
				pingTurn <- struct{}{}
			}
		}
	}()

	pingTurn <- struct{}{}

	sequencia := make([]string, 0, pares*2)
	for i := 0; i < pares*2; i++ {
		sequencia = append(sequencia, <-palavras)
	}
	return sequencia
}
