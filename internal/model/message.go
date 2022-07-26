package model

import "github.com/NganJason/ChatGroup-BE/internal/model/db"

type Message interface {
	GetMessages(channelID *uint64, fromTime *uint64, toTime *uint64, ID *uint64) (messages []*db.Message, err error)
	CreateMessage(req *CreateMessageReq) (message *db.Message, err error)
}

type CreateMessageReq struct {
	MessageID *uint64
	ChannelID *uint64
	UserID    *uint64
	Content   *string
}
