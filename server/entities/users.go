package entities

type User struct {
	ID    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email"`
}

type Users []User
