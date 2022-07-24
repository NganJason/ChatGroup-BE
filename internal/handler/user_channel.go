package handler

import (
	"context"
	"github.com/NganJason/ChatGroup-BE/internal/model"
	"github.com/NganJason/ChatGroup-BE/internal/utils"
	"github.com/NganJason/ChatGroup-BE/vo"
)

type UserChannelHandler struct {
	ctx           context.Context
	userChannelDM model.UserChannel
	channelDM     model.Channel
}

func NewUserChannelHandler(
	ctx context.Context,
	userChannelDM model.UserChannel,
) *UserChannelHandler {
	return &UserChannelHandler{
		ctx:           ctx,
		userChannelDM: userChannelDM,
	}
}

func (h *UserChannelHandler) GetUserChannels(
	userID *uint64,
) (
	voChannels []vo.Channel,
	err error,
) {
	userChannels, err := h.userChannelDM.GetUserChannels(userID, nil)
	if err != nil {
		return nil, err
	}

	channelIDs := make([]*uint64, len(userChannels))
	for i := 0; i < len(channelIDs); i++ {
		channelIDs[i] = userChannels[i].ChannelID
	}

	channels, err := h.channelDM.GetChannels(channelIDs)
	if err != nil {
		return nil, err
	}

	return utils.DBChannelToVo(channels), nil
}
