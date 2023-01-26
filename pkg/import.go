package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/customers"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/invoices"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/products"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/sales"
)

type Utilities interface {
	RunImports() (err error)
	UpdateInvoices() (err error)
}

type utilities struct {
	db *sql.DB
}

func NewUtilities(db *sql.DB) Utilities {
	return &utilities{db: db}
}

func (u *utilities) RunImports() (err error) {
	err = u.ImportProduct()
	if err != nil {
		return
	}
	err = u.ImportCustomers()
	if err != nil {
		return
	}
	err = u.ImportInvoices()
	if err != nil {
		return
	}
	err = u.ImportSales()
	if err != nil {
		return
	}
	return nil
}
func (u *utilities) ImportInvoices() (err error) {
	invoice := []domain.Invoices{}
	data, err := u.getJsonFile("/Users/jvinci/Downloads/desafio-cierre-db-solucion/datos/invoices.json")
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &invoice); err != nil {
		return
	}
	repo := invoices.NewRepository(u.db)
	for _, v := range invoice {
		id, err := repo.Create(&v)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id)
	}
	return nil
}
func (u *utilities) ImportCustomers() (err error) {
	customer := []domain.Customers{}
	data, err := u.getJsonFile("/Users/jvinci/Downloads/desafio-cierre-db-solucion/datos/customers.json")
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &customer); err != nil {
		return
	}
	repo := customers.NewRepository(u.db)
	for _, v := range customer {
		id, err := repo.Create(&v)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id)
	}
	return nil
}
func (u *utilities) ImportProduct() (err error) {
	product := []domain.Product{}
	data, err := u.getJsonFile("/Users/jvinci/Downloads/desafio-cierre-db-solucion/datos/products.json")
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &product); err != nil {
		return
	}
	repo := products.NewRepository(u.db)
	for _, v := range product {
		id, err := repo.Create(&v)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id)
	}
	return nil
}
func (u *utilities) ImportSales() (err error) {
	sale := []domain.Sales{}
	data, err := u.getJsonFile("/Users/jvinci/Downloads/desafio-cierre-db-solucion/datos/sales.json")
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &sale); err != nil {
		return
	}
	repo := sales.NewRepository(u.db)
	for _, v := range sale {
		id, err := repo.Create(&v)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id)
	}
	return nil
}

func (u *utilities) getJsonFile(path string) (data []byte, err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return
	}
	data, err = io.ReadAll(jsonFile)

	if err != nil {
		return
	}
	defer jsonFile.Close()
	return
}

func (u *utilities) UpdateInvoices() (err error) {
	repo := invoices.NewRepository(u.db)
	inv, err := repo.ReadAll()
	if err != nil {
		return
	}
	for _, v := range inv {
		v.Total = repo.CalculateTotal(v.Id)
		v.Datetime = strings.Replace(v.Datetime, "T", " ", 1)
		v.Datetime = strings.Replace(v.Datetime, "Z", "", 1)
		err := repo.Update(v, v.Id)
		if err != nil {
			log.Println(err)
		}
	}
	return
}
