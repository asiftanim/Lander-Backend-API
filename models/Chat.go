package models

import (
	"time"
)

type Chat struct {
	ID         uint
	Message    string
	SenderId   uint
	ReceiverId uint
	Status     uint
	UserType   uint
	CreatedAt  time.Time
}
