package handlers

import (
	"gorm.io/gorm"

	"portfolio/internal/mailer"
)

// Handler bundles the shared dependencies every route handler needs.
type Handler struct {
	DB     *gorm.DB
	Mailer *mailer.Mailer
}

func New(db *gorm.DB, m *mailer.Mailer) *Handler {
	return &Handler{DB: db, Mailer: m}
}
