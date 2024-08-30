package dto

import "time"

type AdminSchedulesRequest struct {
	AdminID     uint      `json:"admin_id"`
	EventName   string    `json:"event_name"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Description string    `json:"description"`
}
