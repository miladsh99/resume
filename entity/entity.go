package entity

import "time"

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type Resume struct {
	ID          int
	User_ID     int
	work_branch string
	Province    string
	City        string
	Address     string
	Zip_code    string
	Birthday    time.Time
	Created_at  time.Time
	Updated_at  time.Time
}

type Notice struct {
	ID          int
	Title       string
	Description string
	Creator_ID  int
	Created_at  time.Time
	Updated_at  time.Time
}
