package models

import (
    "time"
)

type Broker struct{
	ID        	uint 
	Name      	string 
	Email		string 
	Password    string 
	ImagePath  	string
	IsActive    bool  	`json:"is_active"`
	IsDelete    bool	`json:"is_delete"`
	CreatedBy	uint
	CreatedDate time.Time
	ModifiedBy	uint
	ModifiedDate time.Time
}