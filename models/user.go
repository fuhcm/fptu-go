package models

// User ...
type User struct {
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}
