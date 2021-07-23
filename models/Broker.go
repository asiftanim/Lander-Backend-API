package models

import (
	"time"
)

type Broker struct {
	Id         uint
	Name       string
	Email      string
	Password   string
	ImagePath  string
	IsActive   bool
	IsDelete   bool
	CreatedBy  uint
	CreatedAt  time.Time
	ModifiedBy uint
	ModifiedAt time.Time
}
