package response

type UserData struct {
	ID       int    `json:"uid"`
	Username string `json:"username"`
}

type Login struct {
	Token string   `json:"token"`
	User  UserData `json:"user"`
}
