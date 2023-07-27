package dto

type Config struct {
	User struct {
		PasswordSalt   string `json:"password_salt"`
		LoginTokenSalt string `json:"login_token_salt"`
	} `json:"user"`

	Redis struct {
		LoginToken int `json:"login_token"`
	} `json:"redis"`
}
