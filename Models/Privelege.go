package Models

import "time"

//Privilege Structure is Data Privelegies
type Privelege struct {
	Id          int       `json:"id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"` //Omit in case of empty
}

//Privelegios List
type Priveleges []Privelege
