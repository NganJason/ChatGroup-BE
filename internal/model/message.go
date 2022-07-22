package model

import "github.com/NganJason/ChatGroup-BE/internal/model/db"

type Message interface {
	GetMessages(channelID *uint64, fromTime *uint64, toTime *uint64) (messages []db.Message, err error)
}
