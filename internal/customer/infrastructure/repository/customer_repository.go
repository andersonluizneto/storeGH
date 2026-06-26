package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/andersonluizneto/storeGH/internal/customer/domain/entity"
	"github.com/google/uuid"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (r *CustomerRepository) Create(ctx context.Context, customer *entity.Customer) error {

	const query = `
		INSERT INTO cutomers (name, cpf, email, phone) 
		VALUES($1, $2, $3, $4)
		RETURNING id, create_at, update_at	
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		customer.Name,
		customer.Cpf,
		customer.Email,
		customer.Phone,
	).Scan(&customer.ID, &customer.CreatedAt, &customer.UpdatedAt)
}

func (r *CustomerRepository) FindByID(ctx context.Context, id uuid.UUID) (entity.Customer, error) {

	const query = "SELECT id, name, cpf, email, phone FROM customers WHERE id=$1"

	var customer entity.Customer

	err := r.db.QueryRowContext(ctx, query, id).Scan(&customer.ID, &customer.Name, &customer.Cpf, &customer.Email, &customer.Phone)
	if err != nil {
		return customer, err
	}

	return customer, nil

}

func (r *CustomerRepository) FindByCpf(ctx context.Context, cpf string) (entity.Customer, error) {

	const query = "SELECT id, name, cpf, email, phone FROM customers WHERE cpf=$1"

	var customer entity.Customer

	err := r.db.QueryRowContext(ctx, query, cpf).Scan(&customer.ID, &customer.Name, &customer.Cpf, &customer.Email, &customer.Phone)
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *CustomerRepository) FindAll(ctx context.Context) ([]entity.Customer, error) {

	const query = "SELECT id, name, cpf, email, phone FROM customers ORDER BY name"

	rows, err := r.db.QueryContext(ctx, query, nil)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers := make([]entity.Customer, 0)

	for rows.Next() {
		var customer entity.Customer
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.Cpf, &customer.Email, &customer.Phone); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *CustomerRepository) Update(ctx context.Context, customer *entity.Customer) error {

	const query = "UPDATE customers SET name = $1, cpf = $2, email = $3, phone = $4 WHERE id = $5"

	err := r.db.QueryRowContext(
		ctx,
		query,
		customer.Name,
		customer.Cpf,
		customer.Email,
		customer.Phone,
		customer.ID,
	).Scan(&customer.ID)

	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("customer not found")
	}

	return err

}

func (r *CustomerRepository) Delete(ctx context.Context, id uuid.UUID) error {
	const query = "DELETE FROM customers WHERE id = $1"

	var deletedID uuid.UUID
	err := r.db.QueryRowContext(ctx, query, id).Scan(&deletedID)

	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("customer not found")
	}
	return err
}
