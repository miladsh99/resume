package entity

import "time"

type User struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	password  string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetPassword getter
func (u User) GetPassword() string {
	return u.password
}

// SetPassword setter
func (u User) SetPassword(p string) {
	u.password = p
}
