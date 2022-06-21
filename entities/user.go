package entities

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Company struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image"`
}

type Item struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity string `json:"quantity"`
}
