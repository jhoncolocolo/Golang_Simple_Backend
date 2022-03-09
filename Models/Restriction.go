package Models

import "time"

//Restriction Structure is Data Restrictions
type Restriction struct {
	Id          int       `json:"id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"` //Omit in case of empty
}

//Restrictions List
type Restrictions []Restriction
