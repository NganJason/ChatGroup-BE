package model

import "github.com/NganJason/ChatGroup-BE/internal/model/db"

type UserChannel interface {
	GetUserChannels(userID *uint64, channelID *uint64) (userChannels []*db.UserChannel, err error)
	CreateUserChannel(channelID *uint64, userIDs []*uint64) (userChannels []*db.UserChannel, err error)
}
