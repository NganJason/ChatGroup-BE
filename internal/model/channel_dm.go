package model

import (
	"context"

	"github.com/NganJason/ChatGroup-BE/internal/model/db"
)

type ChannelDM struct {
	ctx context.Context
}

func NewChannelDM(ctx context.Context) (Channel, error) {
	return &ChannelDM{
		ctx: ctx,
	}, nil
}

func (dm *ChannelDM) GetChannels(
	channelIDs []*uint64,
) (
	channels []*db.Channel,
	err error,
) {
	return nil, nil
}

func (dm *ChannelDM) GetChannel(
	channelID *uint64,
	channelName *string,
) (
	channel *db.Channel,
	err error,
) {
	return nil, nil
}

func (dm *ChannelDM) CreateChannel(
	req *CreateChannelReq,
) (
	channel *db.Channel,
	err error,
) {
	return nil, nil
}
