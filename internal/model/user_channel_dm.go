package model

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/NganJason/ChatGroup-BE/internal/model/db"
	"github.com/NganJason/ChatGroup-BE/internal/model/query"
	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
)

type UserChannelDM struct {
	ctx context.Context
	db  *sql.DB
}

func NewUserChannelDM(ctx context.Context) (UserChannel, error) {
	return &UserChannelDM{
		ctx: ctx,
	}, nil
}

func (dm *UserChannelDM) GetUserChannels(
	userID *uint64,
	channelID *uint64,
	id *uint64,
) (
	userChannels []*db.UserChannel,
	err error,
) {
	baseQuery := fmt.Sprintf(
		`SELECT * from %s WHERE `,
		dm.getTableName(),
	)

	q := query.NewUserChannelQuery()

	if userID != nil {
		q.UserID(userID)
	}

	if channelID != nil {
		q.ChannelID(channelID)
	}

	if id != nil {
		q.ID(id)
	}

	wheres, args := q.Build()

	rows, err := dm.db.Query(
		baseQuery+wheres,
		args...,
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("query userChannels from db err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	for rows.Next() {
		var userChannel *db.UserChannel

		if err := rows.Scan(
			&userChannel.ID,
			&userChannel.UserID,
			&userChannel.ChannelID,
			&userChannel.Unread,
			&userChannel.CreatedAt,
			&userChannel.UpdatedAt,
		); err != nil {
			if err == sql.ErrNoRows {
				return userChannels, nil
			}

			return nil, cerr.New(
				fmt.Sprintf("query userChannels from db err=%s", err.Error()),
				http.StatusBadGateway,
			)
		}

		userChannels = append(userChannels, userChannel)
	}

	return userChannels, nil
}

func (dm *UserChannelDM) CreateUserChannel(
	channelID *uint64,
	userIDs []*uint64,
) (
	err error,
) {
	q := fmt.Sprintf(
		`
		INSERT INTO %s 
		(user_id, channel_id, unread) 
		VALUES(?, ?, ?)
		`, dm.getTableName(),
	)

	for _, userID := range userIDs {
		_, err := dm.db.Exec(
			q,
			userID,
			channelID,
			0,
		)
		if err != nil {
			return cerr.New(
				fmt.Sprintf("insert userChannel into db err=%s", err.Error()),
				http.StatusBadGateway,
			)
		}
	}

	return nil
}

func (dm *UserChannelDM) getTableName() string {
	return "user_channel_tab"
}
