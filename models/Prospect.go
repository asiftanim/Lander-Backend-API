package models

import (
    "time"
)

type Prospect struct{
	Id        		uint
	FirstName      	string 
	LastName      	string 
	Email			string 
	CookiSecrect    string 
	Company  		string
	IsActive    	bool  	`json:"is_active"`
	IsDelete    	bool	`json:"is_delete"`
	IsEmailVerified   bool	`json:"is_email_verified"`
	CreatedDate 	time.Time
	ModifiedDate 	time.Time
}