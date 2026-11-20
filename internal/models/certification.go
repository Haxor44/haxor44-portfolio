package models

import "time"

// Certification statuses.
const (
	CertEarned     = "earned"
	CertInProgress = "in_progress"
)

// Certification represents a professional certification badge
// (AWS, Kubernetes/CKA, ASCP, etc.), either already earned or in progress.
type Certification struct {
	ID            uint       `gorm:"primaryKey" json:"id"`
	Name          string     `gorm:"size:150;not null" json:"name"`
	IssuingOrg    string     `gorm:"size:150;not null" json:"issuing_org"`
	Status        string     `gorm:"size:20;not null;default:'earned';index" json:"status"` // "earned" or "in_progress"
	IssueDate     *time.Time `json:"issue_date,omitempty"`                                  // nil while in progress
	ExpiryDate    *time.Time `json:"expiry_date,omitempty"`
	CredentialURL string     `gorm:"size:255" json:"credential_url"`
	BadgeIconURL  string     `gorm:"size:255" json:"badge_icon_url"`
	DisplayOrder  int        `gorm:"default:0" json:"display_order"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

// IsInProgress reports whether the certification is still being worked towards.
func (c Certification) IsInProgress() bool {
	return c.Status == CertInProgress
}

// IsExpired reports whether the certification has passed its expiry date.
func (c Certification) IsExpired() bool {
	return c.ExpiryDate != nil && c.ExpiryDate.Before(time.Now())
}
