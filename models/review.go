package models

import "time"

type Review struct {
	ID             uint      `json:"id" db:"id"`
	TenderID       uint      `json:"tender_id" db:"tender_id"`
	BidID          uint      `json:"bid_id" db:"bid_id"`
	Reviewer       string    `json:"reviewer" db:"author_username"`
	Comment        string    `json:"comment" db:"comment"`
	Rating         int       `json:"rating" db:"rating"`
	OrganizationID uint      `json:"organization_id" db:"organization_id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
