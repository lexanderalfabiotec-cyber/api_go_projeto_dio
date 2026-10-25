package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/lexanderalfabiotec-cyber/api_go_projeto_dio/internal/logger"
	"github.com/lexanderalfabiotec-cyber/api_go_projeto_dio/internal/models"
)

func init() {
	// Inicializa o logger para evitar panic nos handlers durante os testes
	logger.Init("error")
}

// newRouter monta um roteador idêntico ao do main, mas usando o handler passado.
func newRouter(h *PersonHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/contato", h.GetPeople).Methods(http.MethodGet)
	r.HandleFunc("/contato/{id}", h.GetPerson).Methods(http.MethodGet)
	r.HandleFunc("/contato/{id}", h.CreatePerson).Methods(http.MethodPost)
	r.HandleFunc("/contato/{id}", h.DeletePerson).Methods(http.MethodDelete)
	return r
}

// ---------- GetPeople ----------

func TestGetPeople_ReturnsAllContacts(t *testing.T) {
	h := NewPersonHandler()
	req := httptest.NewRequest(http.MethodGet, "/contato", nil)
	rr := httptest.NewRecorder()

	newRouter(h).ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("GetPeople: código esperado %d, obteve %d", http.StatusOK, rr.Code)
	}

	var people []models.Person
	if err := json.NewDecoder(rr.Body).Decode(&people); err != nil {
		t.Fatalf("GetPeople: falha ao decodificar resposta: %v", err)
	}
	if len(people) != 2 {
		t.Errorf("GetPeople: esperava 2 contatos, obteve %d", len(people))
	}
}

func TestGetPeople_ContentTypeIsJSON(t *testing.T) {
	h := NewPersonHandler()
	req := httptest.NewRequest(http.MethodGet, "/contato", nil)
	rr := httptest.NewRecorder()

	newRouter(h).ServeHTTP(rr, req)

	ct := rr.Header().Get("Content-Type")
	if ct != "application/json" {
		t.Errorf("GetPeople: Content-Type esperado 'application/json', obteve '%s'", ct)
	}
}

// ---------- GetPerson ----------

func TestGetPerson_ExistingID(t *testing.T) {
	h := NewPersonHandler()
	req := httptest.NewRequest(http.MethodGet, "/contato/1", nil)
	rr := httptest.NewRecorder()

	newRouter(h).ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("GetPerson: código esperado %d, obteve %d", http.StatusOK, rr.Code)
	}

	var p models.Person
	if err := json.NewDecoder(rr.Body).Decode(&p); err != nil {
		t.Fatalf("GetPerson: falha ao decodificar resposta: %v", err)
	}
	if p.ID != "1" {
		t.Errorf("GetPerson: ID esperado '1', obteve '%s'", p.ID)
	}
	if p.Firstname != "John" {
		t.Errorf("GetPerson: Firstname esperado 'John', obteve '%s'", p.Firstname)
	}
}

func TestGetPerson_NotFound(t *testing.T) {
	h := NewPersonHandler()
	req := httptest.NewRequest(http.MethodGet, "/contato/999", nil)
	rr := httptest.NewRecorder()

	newRouter(h).ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("GetPerson: código esperado %d, obteve %d", http.StatusNotFound, rr.Code)
	}
}

// ---------- CreatePerson ----------

func TestCreatePerson_ValidBody(t *testing.T) {
	h := NewPersonHandler()

	newContact := models.Person{
		Firstname: "Alice",
		Lastname:  "Silva",
		Address:   &models.Address{City: "São Paulo", State: "SP"},
	}
	body, _ := json.Marshal(newContact)

	req := httptest.NewRequest(http.MethodPost, "/contato/3", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	newRouter(h).ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("CreatePerson: código esperado %d, obteve %d", http.StatusCreated, rr.Code)
	}

	var people []models.Person
	if err := json.NewDecoder(rr.Body).Decode(&people); err != nil {
		t.Fatalf("CreatePerson: falha ao decodificar resposta: %v", err)
	}
	if len(people) != 3 {
		t.Errorf("CreatePerson: esperava 3 contatos após criação, obteve %d", len(people))
	}

	// Verifica se o ID foi sobrescrito pelo path param
	last := people[len(people)-1]
	if last.ID != "3" {
		t.Errorf("CreatePerson: ID esperado '3', obteve '%s'", last.ID)
	}
}

func TestCreatePerson_InvalidBody(t *testing.T) {
	h := NewPersonHandler()
	req := httptest.NewRequest(http.MethodPost, "/contato/3", bytes.NewReader([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	newRouter(h).ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("CreatePerson body inválido: código esperado %d, obteve %d", http.StatusBadRequest, rr.Code)
	}
}

// ---------- DeletePerson ----------

func TestDeletePerson_ExistingID(t *testing.T) {
	h := NewPersonHandler()
	req := httptest.NewRequest(http.MethodDelete, "/contato/1", nil)
	rr := httptest.NewRecorder()

	newRouter(h).ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("DeletePerson: código esperado %d, obteve %d", http.StatusOK, rr.Code)
	}

	var people []models.Person
	if err := json.NewDecoder(rr.Body).Decode(&people); err != nil {
		t.Fatalf("DeletePerson: falha ao decodificar resposta: %v", err)
	}
	if len(people) != 1 {
		t.Errorf("DeletePerson: esperava 1 contato restante, obteve %d", len(people))
	}
}

func TestDeletePerson_NotFound(t *testing.T) {
	h := NewPersonHandler()
	req := httptest.NewRequest(http.MethodDelete, "/contato/999", nil)
	rr := httptest.NewRecorder()

	newRouter(h).ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("DeletePerson não-encontrado: código esperado %d, obteve %d", http.StatusNotFound, rr.Code)
	}
}
