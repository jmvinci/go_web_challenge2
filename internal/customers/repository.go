package customers

import (
	"database/sql"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

type Repository interface {
	Create(customers *domain.Customers) (int64, error)
	ReadAll() ([]*domain.Customers, error)
	GetTotalAndCondition() (totalAndCondition []domain.TotalAndCondition, err error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(customers *domain.Customers) (int64, error) {
	query := `INSERT INTO customers (first_name, last_name, customers.condition) VALUES (?, ?, ?)`
	row, err := r.db.Exec(query, &customers.FirstName, &customers.LastName, &customers.Condition)
	if err != nil {
		return 0, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) ReadAll() ([]*domain.Customers, error) {
	query := `SELECT id, first_name, last_name, customers.condition FROM customers`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	customers := make([]*domain.Customers, 0)
	for rows.Next() {
		customer := domain.Customers{}
		err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Condition)
		if err != nil {
			return nil, err
		}
		customers = append(customers, &customer)
	}
	return customers, nil
}

func (r *repository) GetTotalAndCondition() (totalAndCondition []domain.TotalAndCondition, err error) {
	query := `SELECT cu.condition as "Condition",  TRUNCATE(SUM(inv.total),2) as Total FROM fantasy_products.invoices inv JOIN fantasy_products.customers cu ON cu.id = inv.customer_id
	GROUP BY cu.condition;`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tac domain.TotalAndCondition
		err := rows.Scan(&tac.Condition, &tac.Total)
		if err != nil {
			return nil, err
		}
		totalAndCondition = append(totalAndCondition, tac)
	}
	return
}
