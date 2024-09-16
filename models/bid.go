package models

import "time"

type Bid struct {
	ID          uint64    `json:"bidId" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Status      string    `json:"status" db:"status"`
	TenderID    uint      `json:"tender_id" db:"tender_id"`
	AuthorID    uint      `json:"author_id" db:"author_id"`
	Version     int64     `json:"version" db:"version"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Decision    string    `json:"decision" db:"decision"`
}
