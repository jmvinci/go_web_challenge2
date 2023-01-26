package domain

type Customers struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Condition bool   `json:"condition"`
}

type TotalAndCondition struct {
	Total     float64 `json:"Total"`
	Condition int     `json:"Condition"`
}
