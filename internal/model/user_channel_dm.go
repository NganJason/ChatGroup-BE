package model

import (
	"context"
	"github.com/NganJason/ChatGroup-BE/internal/model/db"
)

type UserChannelDM struct {
	ctx context.Context
}

func NewUserChannelDM(ctx context.Context) (UserChannel, error) {
	return &UserChannelDM{
		ctx: ctx,
	}, nil
}

func (dm *UserChannelDM) GetUserChannels(
	userID *uint64,
	channelID *uint64,
) (
	userChannels []*db.UserChannel,
	err error,
) {
	return nil, nil
}

func (dm *UserChannelDM) CreateUserChannel(
	channelID *uint64,
	userIDs []*uint64,
	userNames []*string,
) (
	userChannels []*db.UserChannel,
	err error,
) {
	return nil, nil
}
