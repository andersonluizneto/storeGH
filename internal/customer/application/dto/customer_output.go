package dto

import "github.com/google/uuid"

type CustomerOutput struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Cpf   string    `json:"cpf"`
	Email string    `json:"email"`
	Phone string    `json:"phone"`
}
