package model

import "github.com/NganJason/ChatGroup-BE/internal/model/db"

type Message interface {
	GetMessages(channelID *uint64, fromTime *uint64, toTime *uint64) (messages []*db.Message, err error)
	CreateMessage(channelID *uint64, content *string) (message *db.Message, err error)
}
