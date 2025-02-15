package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/bsanzhiev/taco-delivery/common/models"
)

type CustomerRepository interface {
	CreateCustomer(user *models.Customer) (int64, error)
	GetCustomer(id int64) (*models.Customer, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(dsn string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		db: db,
	}, nil
}

func (r *PostgresRepository) CreateCustomer(ctx context.Context, customer *models.Customer) (int64, error) {
	query := `INSERT INTO customers (name, email, phone) VALUES ($1, $2) RETURNING id`
	var id int64
	err := r.db.QueryRow(query, customer.Name, customer.Email, customer.Phone).Scan(&id)
	return id, err
}

func (r *PostgresRepository) GetCustomer(id int64) (*models.Customer, error) {
	query := `SELECT id, name, email, phone FROM customers WHERE id = $1`
	var customer models.Customer
	err := r.db.QueryRow(query, id).Scan(&customer.Id, &customer.Email, &customer.Phone)
	if err == sql.ErrNoRows {
		return nil, errors.New("customer not found")
	}
	return &customer, err
}
