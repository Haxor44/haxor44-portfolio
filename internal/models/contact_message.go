package models

import "time"

// ContactMessage stores a submission from the public contact form.
type ContactMessage struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:150;not null" json:"name"`
	Email     string    `gorm:"size:150;not null" json:"email"`
	Subject   string    `gorm:"size:200" json:"subject"`
	Message   string    `gorm:"type:text;not null" json:"message"`
	IsRead    bool      `gorm:"default:false" json:"is_read"`
	EmailSent bool      `gorm:"default:false" json:"email_sent"`
	CreatedAt time.Time `json:"created_at"`
}
