package handlers

import (
	"fptugo/pkg/core"
	"fptugo/configs/db"
	"fptugo/internal/models"
	"net/http"
)

// ListConfessions ...
func ListConfessions(w http.ResponseWriter, r *http.Request) {
	res := core.Response{ResponseWriter: w}

	confessions := []models.Confession{}
	db.DB.Select(&confessions, "SELECT id, content, sender, approver, reason, created_at, updated_at FROM users")

	res.SendOK(confessions)
}
