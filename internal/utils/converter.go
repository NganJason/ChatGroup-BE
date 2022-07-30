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

func DBUsersToVo(dbUsers []*db.User) []vo.User {
	voUsers := make([]vo.User, len(dbUsers))

	for i := 0; i < len(dbUsers); i++ {
		dbUser := dbUsers[i]

		voUsers[i] = *DBUserToVo(dbUser)
	}

	return voUsers
}

func DBChannelToVo(dbChannel *db.Channel) *vo.Channel {
	return &vo.Channel{
		ChannelID:   dbChannel.ChannelID,
		ChannelName: dbChannel.ChannelName,
		ChannelDesc: dbChannel.ChannelDesc,
	}
}

func DBChannelsToVo(dbChannels []*db.Channel) []vo.Channel {
	voChannels := make([]vo.Channel, len(dbChannels))

	for i := 0; i < len(dbChannels); i++ {
		dbChannel := dbChannels[i]

		voChannels[i] = *DBChannelToVo(dbChannel)
	}

	return voChannels
}

func DBMessageToVo(dbMessage *db.Message, dbUserIDMap map[uint64]*db.User) *vo.Message {
	voMessage := &vo.Message{
		MessageID: dbMessage.MessageID,
		ChannelID: dbMessage.ChannelID,
		Content:   dbMessage.Content,
		CreatedAt: dbMessage.CreatedAt,
	}

	if _, ok := dbUserIDMap[*dbMessage.UserID]; !ok {
		return voMessage
	}

	dbUser := dbUserIDMap[*dbMessage.UserID]
	voUser := DBUserToVo(dbUser)

	voMessage.Sender = voUser

	return voMessage
}

func DBMessagesToVo(dbMessages []*db.Message, dbUserIDMap map[uint64]*db.User) []vo.Message {
	voMessages := make([]vo.Message, len(dbMessages))

	for i := 0; i < len(dbMessages); i++ {
		dbMessage := dbMessages[i]

		voMessages[i] = *DBMessageToVo(dbMessage, dbUserIDMap)
	}

	return voMessages
}
