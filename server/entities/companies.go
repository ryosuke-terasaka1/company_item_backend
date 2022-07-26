package entities

type Company struct {
	ID    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email"`
	Image string `json:"image"`
}

type Companies []Company
