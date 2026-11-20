package models

import "time"

// Project represents a single portfolio project card.
type Project struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	Title              string    `gorm:"size:150;not null" json:"title"`
	Description        string    `gorm:"type:text;not null" json:"description"`
	TechStack          string    `gorm:"size:255;not null" json:"tech_stack"` // comma-separated, e.g. "Go,Fiber,PostgreSQL,Docker"
	GithubURL          string    `gorm:"size:255" json:"github_url"`
	LiveURL            string    `gorm:"size:255" json:"live_url"`
	BlogURL            string    `gorm:"size:255" json:"blog_url"`
	DeploymentPlatform string    `gorm:"size:100" json:"deployment_platform"` // e.g. "AWS ECS", "DigitalOcean", "Railway"
	ImageURL           string    `gorm:"size:255" json:"image_url"`
	Featured           bool      `gorm:"default:false" json:"featured"`
	DisplayOrder       int       `gorm:"default:0" json:"display_order"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

// TechStackList splits the stored comma-separated tech stack into a slice
// for easy rendering as individual badges in templates.
func (p Project) TechStackList() []string {
	return splitAndTrim(p.TechStack)
}
