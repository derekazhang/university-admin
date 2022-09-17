package user

import "time"

type Role struct {
	Type string `json:"role"`
}

type MiniUser struct {
	Id       int       `json:"id"`
	Account  string    `json:"account"`
	Pwd      string    `json:"pwd,omitempty"`
	Name     string    `json:"name"`
	Sex      int       `json:"sex"`
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
	Role     `json:"role"`
}

type UserFull struct {
	Token string `json:"token"`
	MiniUser
}
