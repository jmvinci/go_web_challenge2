package loaddb

type Service interface {
	UpdateInvoices() error
	LoadDb() (err error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) LoadDb() (err error) {
	return s.r.LoadDb()
}

func (s *service) UpdateInvoices() (err error) {
	return s.r.UpdateInvoices()
}
