package models

import (
    "time"
)

type Transaction struct{
	Id        					uint
	ProspectId	     			uint 
	ProspectDomainQueryID    	uint 
	Price   					float64 
	PaymentMethod    			uint 
	CreatedBy					uint
	CreatedDate 				time.Time
	ModifiedBy					uint
	ModifiedDate 				time.Time
	IsActive    				bool  	`json:"is_active"`
	IsDelete    				bool	`json:"is_delete"`
}