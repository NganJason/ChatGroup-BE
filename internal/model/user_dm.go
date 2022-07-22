package model

import (
	"context"

	"github.com/NganJason/ChatGroup-BE/internal/model/db"
)

type UserDM struct {
	ctx context.Context
}

func NewUserDM(ctx context.Context) (User, error) {
	return &UserDM{
		ctx: ctx,
	}, nil
}

func (dm *UserDM) GetUser(
	userID *uint64,
	userName *string,
) (user *db.User, err error) {
	return nil, nil
}

func (dm *UserDM) GetUsers(userID []*uint64) (users []*db.User, err error) {
	return nil, nil
}

func (dm *UserDM) CreateUser(req *CreateUserReq) (user *db.User, err error) {
	return nil, nil
}

func (dm *UserDM) UpdateUser(req *UpdateUserReq) (user *db.User, err error) {
	return nil, nil
}
