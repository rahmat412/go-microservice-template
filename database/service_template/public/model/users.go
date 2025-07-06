package model

import "time"

// Users represents the users table in the database
type Users struct {
	ID        int32      `json:"id"`
	Username  string     `json:"username"`
	FirstName string     `json:"first_name"`
	LastName  *string    `json:"last_name"`
	BirthDate *time.Time `json:"birth_date"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	IsActive  *bool      `json:"is_active"`
	CreatedAt *time.Time `json:"created_at"`
	// Add other fields as needed
}
