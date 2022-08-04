package handler

import (
	"context"

	"github.com/NganJason/ChatGroup-BE/internal/model"
	"github.com/NganJason/ChatGroup-BE/internal/utils"
	"github.com/NganJason/ChatGroup-BE/vo"
)

type UserHandler struct {
	ctx    context.Context
	userDM model.User
}

func NewUserHandler(
	ctx context.Context,
	userDM model.User,
) *UserHandler {
	return &UserHandler{
		ctx:    ctx,
		userDM: userDM,
	}
}

func (h *UserHandler) GetUser(userID *uint64) (userVo *vo.User, err error) {
	user, err := h.userDM.GetUser(userID, nil, nil)
	if err != nil {
		return nil, err
	}

	return utils.DBUserToVo(user), nil
}

func (h *UserHandler) SearchUsers(keyword *string) (usersVo []vo.User, err error) {
	users, err := h.userDM.SearchUsers(keyword)
	if err != nil {
		return nil, err
	}

	return utils.DBUsersToVo(users), nil
}
