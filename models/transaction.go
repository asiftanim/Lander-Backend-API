package models

import (
	"time"
)

type Transaction struct {
	Id                    uint
	ProspectId            uint
	ProspectDomainQueryId uint
	Price                 float64
	PaymentMethod         uint
	CreatedBy             uint
	CreatedAt             time.Time
	ModifiedBy            uint
	ModifiedAt            time.Time
	IsActive              bool
	IsDelete              bool
}
