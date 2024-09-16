package models

import "time"

type Tender struct {
	ID             uint64    `json:"id" db:"id"`
	Name           string    `json:"name" db:"name"`
	Description    string    `json:"description" db:"description"`
	ServiceType    string    `json:"service_type" db:"service_type"`
	Status         string    `json:"status" db:"status"`
	Version        int64     `json:"version" db:"version"`
	OrganizationID uint      `json:"organization_id" db:"organization_id"`
	CreatorID      uint      `json:"creator_id" db:"creator_id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	Reviews        []Review  `json:"reviews,omitempty" db:"-"`
}
