package entity

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID        uuid.UUID
	Name      string
	Cpf       string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
