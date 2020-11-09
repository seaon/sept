package model

type User struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	MobileNo string `json:"mobile_no"`
}