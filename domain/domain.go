package domain

import "time"

type User struct {
	Id         string    `json:"_id"`
	UserName   string    `json:"user_name"`
	NickName   string    `json:"nick_name"`
	Password   string    `json:"password"`
	AvatarUrl  string    `json:"avatar_url"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type Role struct {
	Id         string    `json:"_id"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"create_time"`
}

type Organization struct {
	Id         string    `json:"_id"`
	Name       string    `json:"name"`
	ParentId   string    `json:"parent_id"`
	CreateTime time.Time `json:"create_time"`
}

type UserRole struct {
	Id     string `json:"_id"`
	UserId string `json:"user_id"`
	RoleId string `json:"role_id"`
}

type UserOrg struct {
	Id     string `json:"_id"`
	UserId string `json:"user_id"`
	OrgId  string `json:"org_id"`
}

type Result struct {
	ErrCode    int    `json:"err_code"`
	ErrMessage string `json:"err_message"`
	Data       any    `json:"data"`
}
