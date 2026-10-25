package models

// Person representa um contato na API
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address representa o endereço de um contato
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}
