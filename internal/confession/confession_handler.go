package confession

import (
	"fptugo/configs/db"
	"fptugo/pkg/core"
	"net/http"
)

// ListConfessions ...
func ListConfessions(w http.ResponseWriter, r *http.Request) {
	res := core.Response{ResponseWriter: w}

	confessions := []Confession{}
	db.DB.Select(&confessions, "SELECT id, content, sender, approver, reason, created_at, updated_at FROM users")

	res.SendOK(confessions)
}
