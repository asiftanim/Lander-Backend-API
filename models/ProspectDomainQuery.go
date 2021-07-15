package models

import (
    "time"
)

type ProspectDomainQuery struct{
	Id        			uint
	prospect_id     	uint 
	broker_id      		uint 
	domain_id      		uint 
	offer_type      	uint 
	offered_price   	float64 
	asking_price    	float64 
	Currency			string 
	Status				uint 
	is_price_seen   	bool 
	price_seen_time 	time.Time 
	is_interacted 		bool 
	mail_sent_status	bool 
	mail_sent_time		time.Time  
	CreatedDate 		time.Time
	ModifiedDate 		time.Time
}