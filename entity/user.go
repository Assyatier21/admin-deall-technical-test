package entity

type User struct {
	Id       int64  `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	RoleId   int64  `json:"role_id" form:"role_id"`
	Token    string `json:"token" form:"token"`
}

type UserPoints struct {
	Id     int64 `json:"id" form:"id"`
	Points int64 `json:"points" form:"points"`
}

type RegisteredUser struct {
	Id       int64  `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	RoleId   int64  `json:"role_id" form:"role_id"`
	Token    string `json:"token" form:"token"`
}
