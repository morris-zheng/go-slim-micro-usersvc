package user

import "time"

type User struct {
	Id         int       `json:"id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Name       string    `json:"name"`
}

func (u *User) TableName() string {
	return "user"
}
