package db

import "github.com/NganJason/ChatGroup-BE/internal/vo"

func UserInfoDbToVo(userInfo UserInfo) *vo.UserInfo {
	return &vo.UserInfo{
		UserID:       userInfo.UserID,
		UserName:     userInfo.UserName,
		EmailAddress: userInfo.EmailAddress,
		PhotoURL:     userInfo.PhotoURL,
	}
}
