package models

import "time"

// Confession ...
type Confession struct {
	Content   string    `json:"content" db:"content"`
	Sender    string    `json:"sender" db:"sender"`
	Status    int       `json:"status" db:"status"`
	Approver  string    `json:"approver" db:"approver"`
	Reason    string    `json:"reason" db:"reason"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
