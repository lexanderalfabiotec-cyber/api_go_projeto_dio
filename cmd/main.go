package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lexanderalfabiotec-cyber/api_go_projeto_dio/internal/config"
	"github.com/lexanderalfabiotec-cyber/api_go_projeto_dio/internal/handlers"
	"github.com/lexanderalfabiotec-cyber/api_go_projeto_dio/internal/logger"
)

func main() {
	// 1. Carrega as configurações a partir do config.yml
	cfg, err := config.Load()
	if err != nil {
		// Aviso não-fatal: continua com os valores padrão/env vars
		logger.Init("info")
		logger.Warn("Não foi possível carregar config.yml; usando variáveis de ambiente e valores padrão", "erro", err)
	}

	// 2. Inicializa o logger estruturado
	logger.Init(cfg.LogLevel)

	logger.Info("Iniciando aplicação",
		"app", cfg.AppName,
		"version", cfg.AppVersion,
		"port", cfg.ServerPort,
		"log_level", cfg.LogLevel,
	)

	// 3. Inicializa os handlers
	personHandler := handlers.NewPersonHandler()

	// 4. Configura o roteador e as rotas
	router := mux.NewRouter()

	router.HandleFunc("/contato", personHandler.GetPeople).Methods(http.MethodGet)
	router.HandleFunc("/contato/{id}", personHandler.GetPerson).Methods(http.MethodGet)
	router.HandleFunc("/contato/{id}", personHandler.CreatePerson).Methods(http.MethodPost)
	router.HandleFunc("/contato/{id}", personHandler.DeletePerson).Methods(http.MethodDelete)

	logger.Info("Rotas registradas",
		"GET /contato", "lista todos os contatos",
		"GET /contato/{id}", "busca contato por ID",
		"POST /contato/{id}", "cria novo contato",
		"DELETE /contato/{id}", "remove contato por ID",
	)

	// 5. Sobe o servidor HTTP
	addr := ":" + cfg.ServerPort
	logger.Info("Servidor HTTP iniciado", "endereço", "http://localhost"+addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		logger.Fatal("Falha ao iniciar o servidor HTTP", "error", err)
	}
}
