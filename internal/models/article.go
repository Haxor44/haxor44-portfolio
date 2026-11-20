package models

import "time"

// Article represents a blog post, either linked externally (Medium, dev.to)
// or self-hosted (content stored inline).
type Article struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:200;not null" json:"title"`
	Summary     string    `gorm:"type:text" json:"summary"`
	Content     string    `gorm:"type:text" json:"content,omitempty"` // used only when Platform == "self"
	ExternalURL string    `gorm:"size:255;index" json:"external_url"`
	ImageURL    string    `gorm:"size:255" json:"image_url"`                       // cover image, e.g. from the Medium feed
	Tags        string    `gorm:"size:255" json:"tags"`                            // comma-separated, e.g. "terraform,devops"
	Platform    string    `gorm:"size:50;not null;default:'self'" json:"platform"` // "medium", "devto", "self"
	PublishedAt time.Time `gorm:"not null" json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TagList splits the stored comma-separated tags into a slice for templates.
func (a Article) TagList() []string {
	return splitAndTrim(a.Tags)
}
