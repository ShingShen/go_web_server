package dto

type User struct {
	UserId        uint64  `json:"user_id"`
	UserAccount   string  `json:"user_account"`
	UserPassword  string  `json:"user_password"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	Gender        uint8   `json:"gender"`
	Birthday      string  `json:"birthday"`
	Email         string  `json:"email"`
	Phone         string  `json:"phone"`
	UserProfile   string  `json:"user_profile"`
	Height        *uint16 `json:"height"`
	MedCompliance uint8   `json:"med_compliance"`
	Allergies     *string `json:"allergies"`
	LoginToken    string  `json:"login_token"`
	Role          uint8   `json:"role"`
	Enabled       bool    `json:"enabled"`
	CreatedTime   string  `json:"created_time"`
	UpdatedTime   *string `json:"updated_time"`
}
