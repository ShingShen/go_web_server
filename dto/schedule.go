package dto

type Schedule struct {
	ScheduleId  uint64 `json:"schedule_id"`
	Title       string `json:"title"`
	Note        string `json:"note"`
	StartTime   string `json:"start_time"` // yyyy-mm-dd hh:mm
	EndTime     string `json:"end_time"`   // yyyy-mm-dd hh:mm
	UserId      uint64 `json:"user_id"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}
