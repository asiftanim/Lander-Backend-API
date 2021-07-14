package models

import (
    "time"
)

type Chat struct{
	Id        	uint 
	Message     string 
	SenderId	uint `json:"sender_id"`
	ReceiverId  uint `json:"receiver_id"`
	Status  	uint
	UserType	uint  `json:"user_type"`
	CreatedDate time.Time 
}