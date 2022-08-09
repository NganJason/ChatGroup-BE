package query

import "strings"

type MessageQuery struct {
	channelIDs []*uint64
	fromTime   *uint64
	pageSize   *uint32
	id         *uint64
}

func NewMessageQuery() *MessageQuery {
	return &MessageQuery{}
}

func (q *MessageQuery) ID(
	id *uint64,
) *MessageQuery {
	q.id = id

	return q
}

func (q *MessageQuery) ChannelID(
	channelID *uint64,
) *MessageQuery {
	q.channelIDs = append(q.channelIDs, channelID)

	return q
}

func (q *MessageQuery) ChannelIDs(
	channelIDs []*uint64,
) *MessageQuery {
	for _, id := range channelIDs {
		q.channelIDs = append(q.channelIDs, id)
	}

	return q
}

func (q *MessageQuery) FromTime(
	fromTime *uint64,
) *MessageQuery {
	q.fromTime = fromTime

	return q
}

func (q *MessageQuery) PageSize(
	pageSize *uint32,
) *MessageQuery {
	q.pageSize = pageSize

	return q
}

func (q *MessageQuery) Build() (wheres string, args []interface{}) {
	whereCols := make([]string, 0)

	if q.id != nil {
		whereCols = append(whereCols, "id = ?")
		args = append(args, *q.id)
	}

	if len(q.channelIDs) != 0 {
		inCondition := "channel_id IN (?"

		for i := 1; i < len(q.channelIDs); i++ {
			inCondition = inCondition + ",?"
		}
		inCondition = inCondition + ")"
		whereCols = append(whereCols, inCondition)

		for _, id := range q.channelIDs {
			args = append(args, id)
		}
	}

	if q.fromTime != nil {
		whereCols = append(whereCols, "created_at < ? ")
		args = append(args, q.fromTime)
	}

	wheres = strings.Join(whereCols, " AND ")

	wheres += " ORDER BY created_at DESC"

	if q.pageSize != nil {
		wheres += " LIMIT ?"
		args = append(args, *q.pageSize)
	}

	return wheres, args
}
