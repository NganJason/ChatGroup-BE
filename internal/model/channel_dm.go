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
