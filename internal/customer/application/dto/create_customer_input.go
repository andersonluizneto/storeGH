package dto

type CreateCustomerInput struct {
	Name  string `json:"name"`
	Cpf   string `json:"cpf"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
