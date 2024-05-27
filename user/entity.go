package user

import "time"

type User struct {
	ID           string
	Name         string
	Occupation   string
	Email        string
	PasswordHash string
	Avatar       string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
