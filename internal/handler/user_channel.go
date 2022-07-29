package handler

import (
	"context"
	"net/http"

	"github.com/NganJason/ChatGroup-BE/internal/model"
	"github.com/NganJason/ChatGroup-BE/internal/utils"
	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
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
	userChannels, err := h.userChannelDM.GetUserChannels(userID, nil, nil)
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

	return utils.DBChannelsToVo(channels), nil
}

// Todo: Implement pagination
func (h *UserChannelHandler) GetChannelUsers(
	userID *uint64,
	channelID *uint64,
	pageSize *uint32,
	pageNumber *uint32,
) (
	voUsers []vo.User,
	err error,
) {
	userChannel, err := h.userChannelDM.GetUserChannels(userID, channelID, nil)
	if err != nil {
		return nil, err
	}

	if userChannel == nil {
		return nil, cerr.New(
			"user is not in channel",
			http.StatusBadRequest,
		)
	}

	userChannels, err := h.userChannelDM.GetUserChannels(nil, channelID, nil)
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

func (h *UserChannelHandler) AddUsersToChannel(
	channelID *uint64,
	userIDs []*uint64,
) error {
	err := h.userChannelDM.CreateUserChannel(
		channelID,
		userIDs,
	)
	if err != nil {
		return err
	}

	return nil
}
