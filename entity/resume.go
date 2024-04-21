package entity

import "time"

type Resume struct {
	ID         uint      `json:"id"`
	UserID     uint      `json:"user_id"`
	WorkBranch string    `json:"work_branch"`
	Province   string    `json:"province"`
	City       string    `json:"city"`
	Address    string    `json:"address"`
	ZipCode    string    `json:"zip_code"`
	Birthday   time.Time `json:"birthday"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
