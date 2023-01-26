package invoices

import (
	"database/sql"
	"log"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

type Repository interface {
	Create(invoices *domain.Invoices) (int64, error)
	ReadAll() ([]*domain.Invoices, error)
	CalculateTotal(id int) (total float64)
	Update(invoice *domain.Invoices, id int) (err error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Update(invoice *domain.Invoices, id int) (err error) {
	query := `UPDATE invoices SET datetime=?, customer_id=?, total=? WHERE id = ?`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Println(err)
		err = ErrDataBase
		return
	}
	res, err := stmt.Exec(invoice.Datetime, invoice.CustomerId, invoice.Total, id)
	if err != nil {
		log.Println(err)
		err = ErrDataBase
		return
	}
	_, err = res.RowsAffected()
	if err != nil {
		log.Println(err)
		err = ErrDataBase
		return
	}

	return
}

func (r *repository) CalculateTotal(id int) (total float64) {
	query := `SELECT SUM(sa.quantity * pr.price) as total  FROM fantasy_products.sales sa
	JOIN fantasy_products.products pr ON sa.product_id = pr.id
	JOIN fantasy_products.invoices inv ON sa.invoice_id = inv.id
	WHERE inv.id = ?`

	row := r.db.QueryRow(query, id)
	err := row.Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return
}

func (r *repository) Create(invoices *domain.Invoices) (int64, error) {
	query := `INSERT INTO invoices (customer_id, datetime, total) VALUES (?, ?, ?)`
	row, err := r.db.Exec(query, &invoices.CustomerId, &invoices.Datetime, &invoices.Total)
	if err != nil {
		return 0, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) ReadAll() ([]*domain.Invoices, error) {
	query := `SELECT id, customer_id, datetime, total FROM invoices`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	invoices := make([]*domain.Invoices, 0)
	for rows.Next() {
		invoice := domain.Invoices{}
		err := rows.Scan(&invoice.Id, &invoice.CustomerId, &invoice.Datetime, &invoice.Total)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, &invoice)
	}
	return invoices, nil
}
