package confession

import "time"

// Confession ...
type Confession struct {
	Content   string    `json:"content" db:"content"`
	Sender    string    `json:"sender" db:"sender"`
	Status    int       `json:"status" db:"status"`
	Validator string    `json:"validator" db:"validator"`
	Reason    string    `json:"reason" db:"reason"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
