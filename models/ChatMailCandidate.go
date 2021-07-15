package models

import (
	"time"
)

type ChatMailCandidate struct{
	Id 			uint
	chat_id 	uint
	mail_time 	time.Time
	IsActive    bool  	`json:"is_active"`
}