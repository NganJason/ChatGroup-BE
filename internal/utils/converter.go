package utils

import (
	"github.com/NganJason/ChatGroup-BE/internal/model/db"
	"github.com/NganJason/ChatGroup-BE/vo"
)

func UserDBToVo(userInfo *db.User) *vo.User {
	return &vo.User{
		UserID:       userInfo.UserID,
		UserName:     userInfo.UserName,
		EmailAddress: userInfo.EmailAddress,
		PhotoURL:     userInfo.PhotoURL,
	}
}
