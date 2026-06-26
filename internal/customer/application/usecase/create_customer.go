package usecase

import (
	"context"

	"github.com/andersonluizneto/storeGH/internal/customer/application/dto"
	"github.com/andersonluizneto/storeGH/internal/customer/domain/entity"
	"github.com/andersonluizneto/storeGH/internal/customer/domain/repository"
)

type CreateCustomer struct {
	customerRepository repository.Repository
}

func NewCustomerUseCase(customerRepository repository.Repository) *CreateCustomer {
	return &CreateCustomer{customerRepository: customerRepository}
}

func (cc *CreateCustomer) Execute(ctx context.Context, input dto.CreateCustomerInput) error {
	
	customer := entity.Customer{
		Name: input.Name,
		Cpf: input.Cpf,
		Email: input.Email,
		Phone: input.Phone,
	} 

	return cc.customerRepository.Create(ctx, &customer)
}