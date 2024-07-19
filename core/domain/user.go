package domain

import "context"

type User struct {
	ID          int64  `json:"id"`
	UserName    string `json:"user_name"`
	DisplayName string `json:"display_name"`
	Password    string `json:"password"`
	Avatar      string `json:"avatar"`
	GoogleID    string `json:"google_id"`
	CreateTime  int64  `json:"create_time"`
	UpdateTime  int64  `json:"update_time"`
}

type RepositoryUser interface {
	AddUser(ctx context.Context, req *User) error
	FindUserByUserName(ctx context.Context, userName string) (*User, error)
}
