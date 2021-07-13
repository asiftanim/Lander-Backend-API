package models

type Broker struct{
	Id        	uint 
	Name      	string 
	Email		string 
	Password    string 
	image_url  	string
	IsActive    bool  	`json:"is_active"`
	IsDelete    bool	`json:"is_delete"`
	CreatedBy	uint
	CreatedDate int64
	ModifiedBy	uint
	ModifiedDate int64
}