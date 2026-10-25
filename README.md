# Desafios de lógica e concorrência em Go

Este projeto reúne os desafios de lógica e o exemplo de concorrência desenvolvido em Go.

## Ping-pong com goroutines e canais

O programa principal demonstra concorrência com duas goroutines: uma produz `ping` e a outra produz `pong`. Canais sem buffer fazem a passagem de turno e garantem a alternância determinística:

```text
ping
pong
ping
pong
...
```

## Executar

Na raiz do projeto:

```sh
go run .
```

O programa principal imprime cinco pares. A quantidade pode ser alterada no argumento de `pingpong.GerarSequencia` em [main.go](./main.go).

## Testar

Execute os testes automatizados com:

```sh
go test ./...
```

Os testes verificam a ordem dos pares e o comportamento para quantidades não positivas.

## Desafios de lógica

- `desafio1.ObterDivisiveisPorTres(inicio, fim)` retorna os divisíveis por 3 em um intervalo inclusivo.
- `desafio2.ConverterNumero(n)` retorna `PinPan` para múltiplos de 3 e 5, `Pin` para múltiplos de 3, `Pan` para múltiplos de 5 e o número como texto nos demais casos.

Os testes dos desafios de lógica também são executados por `go test ./...`.

## Como funciona

- A goroutine de `ping` aguarda um sinal em `pingTurn`, envia a palavra e libera `pongTurn`.
- A goroutine de `pong` aguarda `pongTurn`, envia a palavra e libera `pingTurn` enquanto ainda houver pares.
- A função coleta exatamente `2 * pares` mensagens e retorna a sequência para o chamador.
