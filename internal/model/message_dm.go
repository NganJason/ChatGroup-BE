package model

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/NganJason/ChatGroup-BE/internal/config"
	"github.com/NganJason/ChatGroup-BE/internal/model/db"
	"github.com/NganJason/ChatGroup-BE/internal/model/query"
	"github.com/NganJason/ChatGroup-BE/internal/utils"
	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
)

type MessageDM struct {
	ctx context.Context
	db  *sql.DB
}

func NewMessageDM(ctx context.Context) (Message, error) {
	return &MessageDM{
		ctx: ctx,
		db:  config.GetDBs().ChatGroupDB,
	}, nil
}

func (dm *MessageDM) GetMessages(
	channelID *uint64,
	fromTime *uint64,
	toTime *uint64,
	id *uint64,
) (
	messages []*db.Message,
	err error,
) {
	if id == nil {
		if toTime == nil || fromTime == nil {
			return nil, cerr.New(
				"toTime or fromTime cannot be empty",
				http.StatusBadRequest,
			)
		}
	}

	q := query.NewMessageQuery()

	if channelID != nil {
		q.ChannelID(channelID)
	}

	if id != nil {
		q.ID(id)
	}

	if fromTime != nil {
		q.FromTime(fromTime)
	}

	if toTime != nil {
		q.ToTime(toTime)
	}

	wheres, args := q.Build()

	baseQuery := fmt.Sprintf(
		`SELECT * FROM %s WHERE `,
		dm.getTableName(),
	)

	rows, err := dm.db.Query(
		baseQuery+wheres,
		args...,
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("query messages from db err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	for rows.Next() {
		var message db.Message

		if err := rows.Scan(
			&message.ID,
			&message.MessageID,
			&message.ChannelID,
			&message.UserID,
			&message.Content,
			&message.CreatedAt,
			&message.UpdatedAt,
		); err != nil {
			if err == sql.ErrNoRows {
				return messages, nil
			}

			return nil, cerr.New(
				fmt.Sprintf("query messages from db err=%s", err.Error()),
				http.StatusBadGateway,
			)
		}

		messages = append(messages, &message)
	}

	return messages, nil
}

func (dm *MessageDM) CreateMessage(
	req *CreateMessageReq,
) (
	message *db.Message,
	err error,
) {
	q := fmt.Sprintf(
		`
		INSERT INTO %s
		(message_id, channel_id, user_id, content, created_at, updated_at)
		VALUES(?, ?, ?, ?, ?, ?)
		`, dm.getTableName(),
	)

	result, err := dm.db.Exec(
		q,
		req.MessageID,
		req.ChannelID,
		req.UserID,
		req.Content,
		time.Now().UTC().UnixNano(),
		time.Now().UTC().UnixNano(),
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("insert message into db err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	lastInsertID, _ := result.LastInsertId()

	messages, err := dm.GetMessages(
		nil,
		nil,
		nil,
		utils.Uint64Ptr(uint64(lastInsertID)),
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("refetch message from db err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	if len(messages) == 0 {
		return nil, cerr.New(
			"failed to insert message into db",
			http.StatusBadGateway,
		)
	}

	return messages[0], nil
}

func (dm *MessageDM) getTableName() string {
	return "message_tab"
}
