package dto

type Register struct {
	Name         string `json:"name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	PasswordConf string `json:"password_conf"`
	Gender       string `json:"gender"`
}
