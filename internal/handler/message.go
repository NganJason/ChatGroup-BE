package handler

import (
	"context"

	"github.com/NganJason/ChatGroup-BE/internal/model"
	"github.com/NganJason/ChatGroup-BE/internal/model/db"
	"github.com/NganJason/ChatGroup-BE/internal/utils"
	"github.com/NganJason/ChatGroup-BE/vo"
)

type MessageHandler struct {
	ctx       context.Context
	messageDM model.Message
	userDM    model.User
}

func NewMessageHandler(
	ctx context.Context,
	messageDM model.Message,
	userDM model.User,
) *MessageHandler {
	return &MessageHandler{
		ctx:       ctx,
		messageDM: messageDM,
		userDM:    userDM,
	}
}

func (h *MessageHandler) GetMessagesByChannelID(
	channelID,
	fromTime,
	toTime *uint64,
) (
	voMessages []vo.Message,
	err error,
) {
	messages, err := h.messageDM.GetMessages(channelID, fromTime, toTime)
	if err != nil {
		return nil, err
	}

	userIDMap, err := h.userIDMapFromMessages(messages)
	if err != nil {
		return nil, err
	}

	voMessages = utils.DBMessagesToVo(messages, userIDMap)

	return voMessages, nil
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

		dbUser, err := h.userDM.GetUser(userID, nil)
		if err != nil {
			return nil, err
		}

		userIDMap[*userID] = dbUser
	}

	return userIDMap, nil
}