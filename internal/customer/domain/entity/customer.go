package entity

import "github.com/google/uuid"

type Customer struct {
	ID    uuid.UUID
	Name  string
	Cpf   string
	Email string
	Phone string
}
