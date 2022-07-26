package model

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/NganJason/ChatGroup-BE/internal/model/db"
	"github.com/NganJason/ChatGroup-BE/internal/model/query"
	"github.com/NganJason/ChatGroup-BE/internal/utils"
	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
)

type ChannelDM struct {
	ctx context.Context
	db  *sql.DB
}

func NewChannelDM(ctx context.Context) (Channel, error) {
	return &ChannelDM{
		ctx: ctx,
	}, nil
}

func (dm *ChannelDM) GetChannels(
	channelIDs []*uint64,
) (
	channels []*db.Channel,
	err error,
) {
	q := query.NewChannelQuery().ChannelIDs(channelIDs)
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
			fmt.Sprintf("query channels from db err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	for rows.Next() {
		var channel *db.Channel

		if err := rows.Scan(
			&channel.ID,
			&channel.ChannelID,
			&channel.ChannelName,
			&channel.ChannelDesc,
			&channel.Status,
			&channel.CreatedAt,
			&channel.UpdatedAt,
		); err != nil {
			if err == sql.ErrNoRows {
				return channels, nil
			}

			return nil, cerr.New(
				fmt.Sprintf("query channels from db err=%s", err.Error()),
				http.StatusBadGateway,
			)
		}

		channels = append(channels, channel)
	}

	return channels, nil
}

func (dm *ChannelDM) GetChannel(
	channelID *uint64,
	channelName *string,
	id *uint64,
) (
	channel *db.Channel,
	err error,
) {
	q := query.NewChannelQuery()

	if channelID != nil {
		q.ChannelID(channelID)
	}

	if channelName != nil {
		q.ChannelName(channelName)
	}

	if id != nil {
		q.ID(id)
	}

	wheres, args := q.Build()

	baseQuery := fmt.Sprintf(
		`SELECT * from %s WHERE `,
		dm.getTableName(),
	)

	err = dm.db.QueryRow(
		baseQuery+wheres,
		args...,
	).Scan(
		&channel.ID,
		&channel.ChannelID,
		&channel.ChannelName,
		&channel.ChannelDesc,
		&channel.Status,
		&channel.CreatedAt,
		&channel.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, cerr.New(
			fmt.Sprintf("query channel err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	return channel, nil
}

func (dm *ChannelDM) CreateChannel(
	req *CreateChannelReq,
) (
	channel *db.Channel,
	err error,
) {
	q := fmt.Sprintf(
		`
		INSERT INTO %s
		(channel_id, channel_name, channel_desc)
		VALUES(?, ?, ?)
		`, dm.getTableName(),
	)

	result, err := dm.db.Exec(
		q,
		req.ChannelID,
		req.ChannelName,
		req.ChannelDesc,
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("insert channel into db err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	lastInsertID, _ := result.LastInsertId()

	channel, err = dm.GetChannel(
		nil,
		nil,
		utils.Uint64Ptr(uint64(lastInsertID)),
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("refetch channel from db err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	return channel, nil
}

func (dm *ChannelDM) getTableName() string {
	return "channel_tab"
}
