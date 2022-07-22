package model

import "github.com/NganJason/ChatGroup-BE/internal/model/db"

type Channel interface {
	GetChannels(channelIDs []*uint64) (channels []*db.Channel, err error)
}
