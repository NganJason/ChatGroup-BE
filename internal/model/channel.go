package model

import "github.com/NganJason/ChatGroup-BE/internal/model/db"

type Channel interface {
	GetChannel(channelID *uint64, channelName *string) (channel *db.Channel, err error)
	GetChannels(channelIDs []*uint64) (channels []*db.Channel, err error)
	CreateChannel(req *CreateChannelReq) (channel *db.Channel, err error)
}

type CreateChannelReq struct {
	ChannelID   *uint64
	ChannelName *string
	ChannelDesc *string
}
