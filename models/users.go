package models

type User struct {
	ID         string `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	PassWord   string `json:"password" db:"password"`
	Email      string `json:"email" db:"email"`
}
