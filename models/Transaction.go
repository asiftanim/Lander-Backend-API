package models

import (
	"time"
)

type Transaction struct {
	ID                    uint
	ProspectID            uint
	ProspectDomainQueryID uint
	Price                 float64
	PaymentMethod         uint
	CreatedBy             uint
	CreatedAt             time.Time
	ModifiedBy            uint
	ModifiedAt            time.Time
	IsActive              bool `json:"is_active"`
	IsDelete              bool `json:"is_delete"`
}
