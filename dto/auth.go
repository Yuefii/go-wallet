package dto

type Register struct {
	Name         string `json:"name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	PasswordConf string `json:"password_conf"`
	Gender       string `json:"gender"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
