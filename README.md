# Desafios de lógica em Go

Projeto com duas soluções clássicas de lógica de programação: números divisíveis por 3 e a brincadeira PinPan (FizzBuzz).

## Estrutura

```text
.
├── desafio1/
│   ├── divisiveis.go
│   └── divisiveis_test.go
├── desafio2/
│   ├── pinpan.go
│   └── pinpan_test.go
├── go.mod
└── main.go
```

## Como executar

Na raiz do projeto, execute:

```sh
go run .
```

O programa lista os números divisíveis por 3 entre 1 e 100 e, em seguida, exibe o resultado de PinPan para cada número entre 1 e 100.

## Soluções

- `desafio1.ObterDivisiveisPorTres(inicio, fim)` retorna os divisíveis por 3 em um intervalo inclusivo. Um intervalo invertido produz uma fatia vazia.
- `desafio2.ConverterNumero(n)` retorna `PinPan` para múltiplos de 3 e 5, `Pin` para múltiplos de 3, `Pan` para múltiplos de 5 e o número como texto nos demais casos.

## Testes automatizados

Execute os testes de todos os pacotes com:

```sh
go test ./...
```
