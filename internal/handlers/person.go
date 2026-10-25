package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lexanderalfabiotec-cyber/api_go_projeto_dio/internal/logger"
	"github.com/lexanderalfabiotec-cyber/api_go_projeto_dio/internal/models"
)

// PersonHandler agrupa o estado e os métodos dos handlers de Person
type PersonHandler struct {
	people []models.Person
}

// NewPersonHandler cria um novo PersonHandler com dados iniciais de exemplo
func NewPersonHandler() *PersonHandler {
	return &PersonHandler{
		people: []models.Person{
			{
				ID:        "1",
				Firstname: "John",
				Lastname:  "Doe",
				Address:   &models.Address{City: "City X", State: "State X"},
			},
			{
				ID:        "2",
				Firstname: "Koko",
				Lastname:  "Doe",
				Address:   &models.Address{City: "City Z", State: "State Y"},
			},
		},
	}
}

// GetPeople retorna todos os contatos (GET /contato)
func (h *PersonHandler) GetPeople(w http.ResponseWriter, r *http.Request) {
	logger.Info("GetPeople chamado", "method", r.Method, "path", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(h.people); err != nil {
		logger.Error("Erro ao encodar resposta em GetPeople", "error", err)
		http.Error(w, "erro interno do servidor", http.StatusInternalServerError)
	}
}

// GetPerson retorna um único contato pelo ID (GET /contato/{id})
func (h *PersonHandler) GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	logger.Info("GetPerson chamado", "method", r.Method, "path", r.URL.Path, "id", id)
	w.Header().Set("Content-Type", "application/json")

	for _, item := range h.people {
		if item.ID == id {
			if err := json.NewEncoder(w).Encode(item); err != nil {
				logger.Error("Erro ao encodar resposta em GetPerson", "error", err)
				http.Error(w, "erro interno do servidor", http.StatusInternalServerError)
			}
			return
		}
	}

	logger.Warn("Contato não encontrado", "id", id)
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "contato não encontrado", "id": id})
}

// CreatePerson cria um novo contato (POST /contato/{id})
func (h *PersonHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	logger.Info("CreatePerson chamado", "method", r.Method, "path", r.URL.Path, "id", id)
	w.Header().Set("Content-Type", "application/json")

	var person models.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		logger.Error("Erro ao decodificar body em CreatePerson", "error", err)
		http.Error(w, "body inválido", http.StatusBadRequest)
		return
	}

	person.ID = id
	h.people = append(h.people, person)

	logger.Info("Contato criado com sucesso", "id", id)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(h.people); err != nil {
		logger.Error("Erro ao encodar resposta em CreatePerson", "error", err)
	}
}

// DeletePerson remove um contato pelo ID (DELETE /contato/{id})
func (h *PersonHandler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	logger.Info("DeletePerson chamado", "method", r.Method, "path", r.URL.Path, "id", id)
	w.Header().Set("Content-Type", "application/json")

	for index, item := range h.people {
		if item.ID == id {
			h.people = append(h.people[:index], h.people[index+1:]...)
			logger.Info("Contato deletado com sucesso", "id", id)
			if err := json.NewEncoder(w).Encode(h.people); err != nil {
				logger.Error("Erro ao encodar resposta em DeletePerson", "error", err)
			}
			return
		}
	}

	logger.Warn("Tentativa de deletar contato inexistente", "id", id)
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "contato não encontrado", "id": id})
}
