package models

import (
    "time"
)

type Broker struct{
	Id        	uint 
	Name      	string 
	Email		string 
	Password    string 
	image_url  	string
	IsActive    bool  	`json:"is_active"`
	IsDelete    bool	`json:"is_delete"`
	CreatedBy	uint
	CreatedDate time.Time
	ModifiedBy	uint
	ModifiedDate time.Time
}