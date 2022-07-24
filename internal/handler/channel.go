package handler

import (
	"context"
	"net/http"

	"github.com/NganJason/ChatGroup-BE/internal/model"
	"github.com/NganJason/ChatGroup-BE/internal/utils"
	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
	"github.com/NganJason/ChatGroup-BE/vo"
)

type ChannelHandler struct {
	ctx           context.Context
	channelDM     model.Channel
	userChannelDM model.UserChannel
}

func NewChannelHandler(
	ctx context.Context,
	channelDM model.Channel,
) *ChannelHandler {
	return &ChannelHandler{
		ctx:       ctx,
		channelDM: channelDM,
	}
}

func (h *ChannelHandler) SetUserChannelDM(
	userChannelDM model.UserChannel,
) {
	h.userChannelDM = userChannelDM
}

func (h *ChannelHandler) CreateChannel(
	channelName *string,
	channelDesc *string,
	creatorID *uint64,
) (
	*vo.Channel,
	error,
) {
	existingChannel, err := h.channelDM.GetChannel(nil, channelName)
	if err != nil {
		return nil, err
	}

	if existingChannel != nil {
		return nil, cerr.New(
			"channel already exist",
			http.StatusBadRequest,
		)
	}

	channel, err := h.channelDM.CreateChannel(&model.CreateChannelReq{
		ChannelID:   h.generateChannelID(),
		ChannelName: channelName,
		ChannelDesc: channelDesc,
	})
	if err != nil {
		return nil, err
	}

	_, err = h.userChannelDM.CreateUserChannel(
		channel.ChannelID,
		[]*uint64{
			creatorID,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}

	return utils.DBChannelToVo(channel), nil
}

func (h *ChannelHandler) generateChannelID() *uint64 {
	return nil
}
