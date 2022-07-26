package model

import "github.com/NganJason/ChatGroup-BE/internal/model/db"

type User interface {
	GetUser(userID *uint64, userName *string, id *uint64) (user *db.User, err error)
	GetUsers(userID []*uint64) (users []*db.User, err error)
	CreateUser(req *CreateUserReq) (user *db.User, err error)
	UpdateUser(req *UpdateUserReq) (user *db.User, err error)
}

type CreateUserReq struct {
	UserID         *uint64
	UserName       string
	HashedPassword string
	Salt           string
	EmailAddress   *string
	PhotoURL       *string
}

type UpdateUserReq struct {
	UserID         uint64
	UserName       *string
	HashedPassword *string
	Salt           *string
	EmailAddress   *string
	PhotoURL       *string
}
