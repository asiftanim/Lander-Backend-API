package models

import (
	"time"
)

type ChatMailCandidate struct {
	ID        uint
	ChatId    uint
	MailTime  time.Time
	IsActive  bool
	CreatedAt time.Time
}
