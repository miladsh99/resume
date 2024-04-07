package entity

import "time"

type Resume struct {
	ID         uint
	UserID     uint64
	WorkBranch string
	Province   string
	City       string
	Address    string
	ZipCode    string
	Birthday   time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Notice struct {
	ID          uint
	Title       string
	Description string
	CreatorID   uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
