package handler

import (
	"context"
	"net/http"

	"github.com/NganJason/ChatGroup-BE/internal/model"
	"github.com/NganJason/ChatGroup-BE/internal/utils"
	"github.com/NganJason/ChatGroup-BE/pkg/auth"
	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
	"github.com/NganJason/ChatGroup-BE/vo"
)

type AuthHandler struct {
	ctx    context.Context
	userDM model.User
}

func NewAuthHandler(
	ctx context.Context,
	userDM model.User,
) *AuthHandler {
	return &AuthHandler{
		ctx:    ctx,
		userDM: userDM,
	}
}

func (h *AuthHandler) Login(
	userName *string,
	password *string,
) (userVo *vo.User, err error) {
	user, err := h.userDM.GetUser(
		nil,
		userName,
	)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, cerr.New(
			"user not found",
			http.StatusBadRequest,
		)
	}

	isPasswordMatch := auth.ComparePasswordSHA(
		*password,
		*user.HashedPassword,
		*user.Salt,
	)

	if !isPasswordMatch {
		return nil, cerr.New(
			"invalid password",
			http.StatusUnauthorized,
		)
	}

	return utils.UserDBToVo(user), nil
}

func (h *AuthHandler) Signup(
	userName *string,
	password *string,
) (userVo *vo.User, err error) {
	existingUser, err := h.userDM.GetUser(
		nil,
		userName,
	)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, cerr.New(
			"user already exist",
			http.StatusBadRequest,
		)
	}

	hashedPassword, salt := auth.CreatePasswordSHA(
		*password,
		16,
	)

	req := &model.CreateUserReq{
		UserName:       *userName,
		HashedPassword: hashedPassword,
		Salt:           salt,
	}

	newUser, err := h.userDM.CreateUser(req)
	if err != nil {
		return nil, err
	}

	return utils.UserDBToVo(newUser), nil
}
