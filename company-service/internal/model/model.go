package model

import (
	"github.com/google/uuid"
)

type Company struct {
	ID			uuid.UUID	`grom:"primaryKey" json:"id"`
	Name		string		`grom:"type:varchar(15)" json:"name"`
	Description string		`grom:"type:varchar(3000)" json:"description"`
	Employees	int			`grom:"unique" json:"employees"`
	Registered 	bool		`json:"registered"`
	Type		string		`json:"type"`
}