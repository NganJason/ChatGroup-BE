package model

import "github.com/NganJason/ChatGroup-BE/internal/model/db"

type UserModel interface {
	GetUserInfo(userID *uint64, userName *uint64) (userInfo *db.UserInfo, err error)
	GetUserInfos(userID []*uint64) (userInfos []*db.UserInfo, err error)
}
