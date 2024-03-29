package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/ChatGroup-BE/internal/model"
	"github.com/NganJason/ChatGroup-BE/internal/model/db"
	"github.com/NganJason/ChatGroup-BE/internal/utils"
	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
	"github.com/NganJason/ChatGroup-BE/pkg/clog"
	"github.com/NganJason/ChatGroup-BE/pkg/socket"
	"github.com/NganJason/ChatGroup-BE/vo"
)

type MessageHandler struct {
	ctx           context.Context
	messageDM     model.Message
	userDM        model.User
	userChannelDM model.UserChannel
}

func NewMessageHandler(
	ctx context.Context,
	messageDM model.Message,
	userDM model.User,
	userChannelDM model.UserChannel,
) *MessageHandler {
	return &MessageHandler{
		ctx:           ctx,
		messageDM:     messageDM,
		userDM:        userDM,
		userChannelDM: userChannelDM,
	}
}

func (h *MessageHandler) GetMessagesByChannelID(
	userID,
	channelID,
	fromTime *uint64,
	pageSize *uint32,
) (
	voMessages []vo.Message,
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

	messages, err := h.messageDM.GetMessages(
		channelID,
		fromTime,
		pageSize,
		nil,
	)
	if err != nil {
		return nil, err
	}

	userIDMap, err := h.userIDMapFromMessages(messages)
	if err != nil {
		return nil, err
	}

	voMessages = utils.DBMessagesToVo(
		messages,
		userIDMap,
	)

	return voMessages, nil
}

func (h *MessageHandler) CreateMessage(
	channelID *uint64,
	content *string,
	userID *uint64,
) (
	*vo.Message,
	error,
) {
	userChannel, err := h.userChannelDM.GetUserChannels(
		userID,
		channelID,
		nil,
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("get existing userChannel err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	if userChannel == nil {
		return nil, cerr.New(
			"creator is not in channel/ channel does not exist",
			http.StatusBadRequest,
		)
	}

	message, err := h.messageDM.CreateMessage(
		&model.CreateMessageReq{
			MessageID: utils.GenerateUUID(),
			ChannelID: channelID,
			Content:   content,
			UserID:    userID,
		},
	)
	if err != nil {
		return nil, err
	}

	userIDMap, err := h.userIDMapFromMessages([]*db.Message{message})
	if err != nil {
		return nil, err
	}

	return utils.DBMessageToVo(message, userIDMap), nil
}

func (h *MessageHandler) BroadcastMessage(
	hub socket.Hub,
	senderID *uint64,
	channelID *uint64,
	message *vo.Message,
) {
	if hub == nil {
		clog.Error(
			h.ctx,
			"hub is nil in broadcastMessage",
		)
		return
	}

	userChannels, err := h.userChannelDM.GetUserChannels(
		nil,
		channelID,
		nil,
	)
	if err != nil {
		clog.Error(
			h.ctx,
			fmt.Sprintf("get userChannel in broadcastMsg err=%s", err.Error()),
		)
		return
	}

	recipientIDs := make([]uint64, 0)
	for _, d := range userChannels {
		if d.UserID == nil {
			continue
		}

		if *d.UserID == *senderID {
			continue
		}

		recipientIDs = append(recipientIDs, *d.UserID)
	}

	hub.Broadcast(
		*senderID,
		recipientIDs,
		vo.SocketMessage{
			EventType: uint32(vo.ClientSendMsgEvent),
			Message:   message,
		},
	)
}

func (h *MessageHandler) userIDMapFromMessages(messages []*db.Message) (map[uint64]*db.User, error) {
	userIDMap := make(map[uint64]*db.User)

	for _, message := range messages {
		userID := message.UserID
		if userID == nil {
			continue
		}

		if _, ok := userIDMap[*userID]; ok {
			continue
		}

		dbUser, err := h.userDM.GetUser(userID, nil, nil)
		if err != nil {
			return nil, err
		}

		userIDMap[*userID] = dbUser
	}

	return userIDMap, nil
}
