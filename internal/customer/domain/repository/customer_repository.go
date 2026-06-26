package repository

import (
	"context"

	"github.com/andersonluizneto/storeGH/internal/customer/domain/entity"
	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, customer *entity.Customer) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Customer, error)
	FindByCpf(ctx context.Context, cpf string) (*entity.Customer, error)
	FindAll(ctx context.Context) ([]*entity.Customer, error)
	Update(ctx context.Context, customer *entity.Customer, id uuid.UUID) (uint8, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
