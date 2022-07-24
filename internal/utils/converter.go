package utils

import (
	"github.com/NganJason/ChatGroup-BE/internal/model/db"
	"github.com/NganJason/ChatGroup-BE/vo"
)

func DBUserToVo(dbUser *db.User) *vo.User {
	return &vo.User{
		UserID:       dbUser.UserID,
		UserName:     dbUser.UserName,
		EmailAddress: dbUser.EmailAddress,
		PhotoURL:     dbUser.PhotoURL,
	}
}

func DBChannelToVo(dbChannels []*db.Channel) []vo.Channel {
	voChannels := make([]vo.Channel, len(dbChannels))

	for i := 0; i < len(dbChannels); i++ {
		dbChannel := dbChannels[i]

		voChannels[i] = vo.Channel{
			ChannelID:   dbChannel.ChannelID,
			ChannelName: dbChannel.ChannelName,
			ChannelDesc: dbChannel.ChannelDesc,
		}
	}

	return voChannels
}

func DBMessagesToVo(dbMessages []*db.Message, dbUserIDMap map[uint64]*db.User) []vo.Message {
	voMessages := make([]vo.Message, len(dbMessages))

	for i := 0; i < len(dbMessages); i++ {
		dbMessage := dbMessages[i]

		voMessages[i] = vo.Message{
			MessageID: dbMessage.MessageID,
			ChannelID: dbMessage.ChannelID,
			Content:   dbMessage.Content,
			CreatedAt: dbMessage.CreatedAt,
		}

		if _, ok := dbUserIDMap[*dbMessage.UserID]; !ok {
			continue
		}

		dbUser := dbUserIDMap[*dbMessage.UserID]
		voUser := DBUserToVo(dbUser)

		voMessages[i].Sender = voUser
	}

	return voMessages
}
