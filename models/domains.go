package models

import (
	"time"
)

type Domain struct {
	Id          uint
	DomainId    uint
	Name        string
	PriceEur    float64
	PriceUsd    float64
	Company     string
	IsPurshable bool
	ModifiedAt  time.Time
}
