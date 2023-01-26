package loaddb

import (
	"database/sql"

	"github.com/bootcamp-go/desafio-cierre-db.git/pkg"
)

type Repository interface {
	UpdateInvoices() (err error)
	LoadDb() (err error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) UpdateInvoices() (err error) {
	imports := pkg.NewUtilities(r.db)
	err = imports.UpdateInvoices()
	if err != nil {
		return
	}
	return
}

func (r *repository) LoadDb() (err error) {
	imports := pkg.NewUtilities(r.db)
	err = imports.RunImports()
	if err != nil {
		return
	}
	return
}
