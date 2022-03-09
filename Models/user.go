package Models

import (
	"time"
)

//User datos del usuario
type User struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

//Users lista de usuarios
type Users []*User
