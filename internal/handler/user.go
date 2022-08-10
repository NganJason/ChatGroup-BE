package handler

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/NganJason/ChatGroup-BE/internal/model"
	"github.com/NganJason/ChatGroup-BE/internal/utils"
	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
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

func (h *UserHandler) UploadImage(
	fileByte []byte,
	userID uint64,
) (string, error) {
	if exist := utils.IsDirExist(utils.ImageDir); !exist {
		if err := os.Mkdir(utils.ImageDir, os.ModePerm); err != nil {
			return "", cerr.New(
				fmt.Sprintf("mkdir err=%s", err.Error()),
				http.StatusBadGateway,
			)
		}
	}

	tempFile, err := ioutil.TempFile(utils.ImageDir, "*.png")
	if err != nil {
		return "", cerr.New(
			fmt.Sprintf("create file err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}
	defer tempFile.Close()

	tempFile.Write(fileByte)
	fileName := tempFile.Name()

	user, err := h.userDM.UpdateUser(
		&model.UpdateUserReq{
			UserID:   userID,
			PhotoURL: utils.StrPtr(fileName),
		},
	)
	if err != nil {
		return "", err
	}

	return *user.PhotoURL, nil
}
