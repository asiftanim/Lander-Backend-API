package models

import (
	"time"
)

type ProspectDomainQuery struct {
	Id             uint
	ProspectId     uint
	BrokerId       uint
	DomainId       uint
	OfferType      uint
	OfferedPrice   float64
	AskingPrice    float64
	Currency       string
	Status         uint
	IsPriceSeen    bool
	PriceSeenTime  time.Time
	IsInteracted   bool
	MailSentStatus bool
	MailSentTime   time.Time
	CreatedAt      time.Time
	ModifiedAt     time.Time
}
