package models

// User ...
type User struct {
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Level    int    `json:"level" db:"level"`
}
