package invoices

import (
	"errors"
	"fmt"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

var (
	ErrNotFound       = errors.New("error: product not found")
	ErrDuplicatedCode = errors.New("error: duplicated code")
	ErrInvalidFormat  = errors.New("error: invalid format")
	ErrInternal       = errors.New("error: server error")
	ErrDataBase       = errors.New("error: database error")
	ErrCodeNotFound   = errors.New("error: code not found")
)

type Service interface {
	Create(invoices *domain.Invoices) error
	ReadAll() ([]*domain.Invoices, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(invoices *domain.Invoices) error {
	_, err := s.r.Create(invoices)
	if err != nil {
		return err
	}
	return nil

}
func (s *service) ReadAll() ([]*domain.Invoices, error) {
	fmt.Println(s.r.CalculateTotal(4))
	return s.r.ReadAll()
}
