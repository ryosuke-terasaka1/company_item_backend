package entities

type Item struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"not null"`
	Price    int    `json:"price"`
	Quantity string `json:"quantity"`
}

type Items []Item
