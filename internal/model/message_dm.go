package model

import (
	"context"

	"github.com/NganJason/ChatGroup-BE/internal/model/db"
)

type MessageDM struct {
	ctx context.Context
}

func NewMessageDM(ctx context.Context) (Message, error) {
	return &MessageDM{
		ctx: ctx,
	}, nil
}

func (dm *MessageDM) GetMessages(
	channelID *uint64,
	fromTime *uint64,
	toTime *uint64,
) (
	messages []*db.Message,
	err error,
) {
	return nil, nil
}

func (dm *MessageDM) CreateMessage(
	channelID *uint64,
	content *string,
) (
	message *db.Message,
	err error,
) {
	return nil, nil
}
