package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NganJason/ChatGroup-BE/internal/model"
	"github.com/NganJason/ChatGroup-BE/internal/model/db"
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
	creatorID *uint64,
	channelID *uint64,
	userIDs []*uint64,
) error {
	userChannel, err := h.userChannelDM.GetUserChannels(
		creatorID,
		channelID,
		nil,
	)
	if err != nil {
		return cerr.New(
			fmt.Sprintf("get existing userChannel err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	if userChannel == nil {
		return cerr.New(
			"creator is not in channel/ channel does not exist",
			http.StatusBadRequest,
		)
	}

	existingUsers, err := h.userDM.GetUsers(userIDs)
	if err != nil {
		return cerr.New(
			fmt.Sprintf("get existing users err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	if len(existingUsers) != len(userIDs) {
		invalidUserIDs := h.getInvalidUserIDs(existingUsers, userIDs)

		msg := "invalid users="
		for _, id := range invalidUserIDs {
			msg = msg + strconv.Itoa(int(id)) + ","
		}

		return cerr.New(
			msg,
			http.StatusBadRequest,
		)
	}

	err = h.userChannelDM.CreateUserChannel(
		channelID,
		userIDs,
	)
	if err != nil {
		return err
	}

	return nil
}

func (h *UserChannelHandler) getInvalidUserIDs(
	existingUsers []*db.User,
	userIDs []*uint64,
) []uint64 {
	userIDMap := make(map[uint64]bool)
	invalidUserIDs := make([]uint64, 0)

	for _, user := range existingUsers {
		userIDMap[*user.UserID] = true
	}

	for _, id := range userIDs {
		if _, ok := userIDMap[*id]; !ok {
			invalidUserIDs = append(invalidUserIDs, *id)
		}
	}

	return invalidUserIDs
}
