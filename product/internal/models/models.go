package models

import (
	"time"
)

type Product struct {
	ID          int64
	Name        string
	Description string
	Price       float64
	Category    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Review struct {
	ID        int64
	ProductID int64
	UserID    int64
	Rating    int
	Comment   string
	CreatedAt time.Time
}
