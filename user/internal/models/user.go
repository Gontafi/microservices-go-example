package models

import (
	"time"
)

type User struct {
	ID        int64
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserProfile struct {
	ID        int64
	UserID    int64
	FirstName string
	LastName  string
}
