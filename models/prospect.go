package models

import (
	"time"
)

type Prospect struct {
	Id              uint
	FirstName       string
	LastName        string
	Email           string
	CookiSecrect    string
	Company         string
	IsActive        bool
	IsDelete        bool
	IsEmailVerified bool
	CreatedAt       time.Time
	ModifiedAt      time.Time
}
