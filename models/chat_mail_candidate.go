package models

import (
	"time"
)

type ChatMailCandidate struct {
	Id        uint
	ChatId    uint
	MailTime  time.Time
	IsActive  bool
	CreatedAt time.Time
}
