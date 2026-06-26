package handler

import (
	"net/http"

	"github.com/andersonluizneto/storeGH/internal/customer/application/usecase"
)

type CustomerHandler struct {
	createCustomer usecase.CreateCustomer
}

func NewCustomerUseCase(createCustomer usecase.CreateCustomer) *CustomerHandler {
	return &CustomerHandler{createCustomer: createCustomer}
}

func (ch *CustomerHandler) Create(w http.ResponseWriter, r *http.Request)  {
}

func (ch *CustomerHandler) FindAll(w http.ResponseWriter, r *http.Request)  {
}

func (ch *CustomerHandler) FindByID(w http.ResponseWriter, r *http.Request)  {
}

func (ch *CustomerHandler) Update(w http.ResponseWriter, r *http.Request)  {
}

func (ch *CustomerHandler) Delete(w http.ResponseWriter, r *http.Request)  {
}
