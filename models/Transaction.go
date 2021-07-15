package models

import (
    "time"
)

type Transaction struct{
	Id        					uint
	prospect_id     			uint 
	prospect_domain_query_id    uint 
	Price   					float64 
	payment_method    			uint 
	CreatedBy					uint
	CreatedDate 				time.Time
	ModifiedBy					uint
	ModifiedDate 				time.Time
	IsActive    				bool  	`json:"is_active"`
	IsDelete    				bool	`json:"is_delete"`
}