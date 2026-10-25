# API de Contatos — Projeto DIO (Go)

API REST desenvolvida em Go como desafio do bootcamp da [DIO](https://www.dio.me/), com configuração via arquivo YAML, logger estruturado (JSON) e testes unitários.

---

## Sumário

- [Visão Geral](#visão-geral)
- [Tecnologias](#tecnologias)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Configuração](#configuração)
- [Executando a API](#executando-a-api)
- [Endpoints](#endpoints)
- [Testes](#testes)

---

## Visão Geral

A API gerencia uma lista de **contatos** (`Person`), expondo operações de listagem, consulta por ID, criação e remoção. Os dados são mantidos em memória — sem banco de dados externo.

---

## Tecnologias

| Tecnologia | Versão |
|---|---|
| Go | 1.27+ |
| [gorilla/mux](https://github.com/gorilla/mux) | v1.8.1 |
| [gopkg.in/yaml.v3](https://pkg.go.dev/gopkg.in/yaml.v3) | v3.0.1 |
| `log/slog` (stdlib) | — |

---

## Estrutura do Projeto

```
api_go_projeto_dio/
├── cmd/
│   └── main.go                   # Ponto de entrada da aplicação
├── internal/
│   ├── config/
│   │   ├── config.go             # Carregamento de configurações (YAML + env vars)
│   │   └── config_test.go        # Testes unitários do pacote config
│   ├── handlers/
│   │   ├── person.go             # Handlers HTTP dos endpoints de contato
│   │   └── person_test.go        # Testes unitários dos handlers
│   ├── logger/
│   │   └── logger.go             # Logger estruturado (JSON) baseado em slog
│   └── models/
│       └── person.go             # Modelos Person e Address
├── config.yml                    # Arquivo de configuração da aplicação
├── go.mod
└── go.sum
```

---

## Configuração

As configurações são carregadas em **cascata**:

1. **`config.yml`** (prioridade máxima)
2. **Variáveis de ambiente**
3. **Valores padrão**

### `config.yml`

```yaml
server:
  port: "8000"       # Porta HTTP do servidor

log:
  level: "info"      # Nível de log: debug | info | warn | error

app:
  name: "contatos-api"
  version: "1.0.0"
```

### Variáveis de ambiente equivalentes

| Variável | Padrão | Descrição |
|---|---|---|
| `SERVER_PORT` | `8000` | Porta do servidor HTTP |
| `LOG_LEVEL` | `info` | Nível de log |
| `APP_NAME` | `contatos-api` | Nome da aplicação |
| `APP_VERSION` | `1.0.0` | Versão da aplicação |
| `CONFIG_PATH` | `config.yml` | Caminho alternativo para o arquivo YAML |

---

## Executando a API

```bash
# Clonar o repositório
git clone https://github.com/lexanderalfabiotec-cyber/api_go_projeto_dio.git
cd api_go_projeto_dio

# Instalar dependências
go mod download

# Executar
go run ./cmd/main.go
```

O servidor sobe em `http://localhost:8000` (ou na porta definida em `config.yml`).

---

## Endpoints

Base URL: `http://localhost:8000`

| Método | Rota | Descrição |
|---|---|---|
| `GET` | `/contato` | Lista todos os contatos |
| `GET` | `/contato/{id}` | Retorna um contato pelo ID |
| `POST` | `/contato/{id}` | Cria um novo contato com o ID informado |
| `DELETE` | `/contato/{id}` | Remove um contato pelo ID |

### Modelo `Person`

```json
{
  "id": "1",
  "firstname": "John",
  "lastname": "Doe",
  "address": {
    "city": "City X",
    "state": "State X"
  }
}
```

### Exemplos com `curl`

```bash
# Listar todos os contatos
curl http://localhost:8000/contato

# Buscar contato por ID
curl http://localhost:8000/contato/1

# Criar novo contato
curl -X POST http://localhost:8000/contato/3 \
  -H "Content-Type: application/json" \
  -d '{"firstname":"Alice","lastname":"Silva","address":{"city":"São Paulo","state":"SP"}}'

# Deletar contato
curl -X DELETE http://localhost:8000/contato/1
```

### Respostas de Erro

| Código | Situação |
|---|---|
| `400 Bad Request` | Body inválido na criação |
| `404 Not Found` | ID não encontrado |
| `500 Internal Server Error` | Falha ao serializar resposta |

---

## Testes

```bash
# Executar todos os testes
go test ./...

# Com saída detalhada
go test ./... -v

# Com cobertura
go test ./... -cover
```

### Cobertura atual

| Pacote | Testes |
|---|---|
| `internal/config` | `loadYAML`, `firstNonEmpty`, `getEnv`, `Load` (padrão e env vars) |
| `internal/handlers` | `GetPeople`, `GetPerson`, `CreatePerson`, `DeletePerson` — sucesso e erros |
