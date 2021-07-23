package models

import (
	"time"
)

type Chat struct {
	Id         uint
	Message    string
	SenderId   uint
	ReceiverId uint
	Status     uint
	UserType   uint
	CreatedAt  time.Time
}
