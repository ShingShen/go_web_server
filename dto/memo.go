package dto

type Memo struct {
	MemoId      uint64  `json:"memo_id"`
	Content     string  `json:"content"`
	HasRead     uint8   `json:"has_read"`
	CreatedBy   *string `json:"created_by"`
	UpdatedBy   *string `json:"updated_by"`
	UserId      uint64  `json:"user_id"`
	CreatedTime string  `json:"created_time"`
	UpdatedTime *string `json:"updated_time"`
}
