# Conversor de escalas de temperatura

Este projeto demonstra como converter temperaturas entre as escalas Celsius,
Kelvin e Fahrenheit usando funções em Go.

## Conversões disponíveis

O programa possui uma função para cada conversão possível:

- Celsius para Kelvin: `K = C + 273`
- Kelvin para Celsius: `C = K - 273`
- Celsius para Fahrenheit: `F = (C × 9 / 5) + 32`
- Fahrenheit para Celsius: `C = (F - 32) × 5 / 9`
- Kelvin para Fahrenheit: primeiro converte para Celsius e depois para Fahrenheit
- Fahrenheit para Kelvin: primeiro converte para Celsius e depois para Kelvin

## Demonstração

Ao executar o programa, são demonstradas todas as conversões usando valores
equivalentes a 25 °C, 298 K e 77 °F:

```text
Conversões de temperatura

25.00 °C = 298.00 K
25.00 °C = 77.00 °F
298.00 K = 25.00 °C
298.00 K = 77.00 °F
77.00 °F = 25.00 °C
77.00 °F = 298.00 K
```

## Como executar

É necessário ter o Go instalado. No terminal, dentro da pasta do projeto,
execute:

```bash
go run main.go
```

## Observação

Neste exercício foi utilizada a aproximação `273` para a conversão entre
Celsius e Kelvin, conforme a fórmula proposta no enunciado. Em aplicações
científicas, normalmente utiliza-se `273.15`.
