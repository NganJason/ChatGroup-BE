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
	userDM        model.User
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

func (h *UserChannelHandler) SetChannelDM(
	channelDM model.Channel,
) {
	h.channelDM = channelDM
}

func (h *UserChannelHandler) SetUserDM(
	userDM model.User,
) {
	h.userDM = userDM
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

// Todo: Implement pagination
func (h *UserChannelHandler) GetChannelUsers(
	channelID *uint64,
	pageSize *uint32,
	pageNumber *uint32,
) (
	voUsers []vo.User,
	err error,
) {
	userChannels, err := h.userChannelDM.GetUserChannels(nil, channelID)
	if err != nil {
		return nil, err
	}

	userIDs := make([]*uint64, len(userChannels))
	for i := 0; i < len(userIDs); i++ {
		userIDs[i] = userChannels[i].UserID
	}

	users, err := h.userDM.GetUsers(userIDs)
	if err != nil {
		return nil, err
	}

	return utils.DBUsersToVo(users), nil
}
