# Calculadora em Go

Projeto simples de calculadora com as operações básicas de soma, subtração,
multiplicação e divisão.

## Operações

- `Soma(a, b)`: retorna `a + b`.
- `Subtrai(a, b)`: retorna `a - b`.
- `Multiplica(a, b)`: retorna `a * b`.
- `Divide(a, b)`: retorna `a / b` e um erro se `b` for zero.

Os argumentos e resultados usam `float64`, permitindo trabalhar com números
inteiros e decimais.

## Executar os testes

Na raiz do projeto, execute:

```sh
go test ./...
```

Para ver também cada teste executado:

```sh
go test -v ./...
```
