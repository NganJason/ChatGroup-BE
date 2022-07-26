package query

import "strings"

type MessageQuery struct {
	channelIDs []*uint64
	fromTime   *uint64
	toTime     *uint64
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

func (q *MessageQuery) ToTime(
	toTime *uint64,
) *MessageQuery {
	q.toTime = toTime

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

	if q.fromTime != nil && q.toTime != nil {
		betweenCondition := "created_at BETWEEN ? AND ? "
		whereCols = append(whereCols, betweenCondition)
		args = append(args, q.fromTime)
		args = append(args, q.toTime)
	}

	wheres = strings.Join(whereCols, " AND ")

	return wheres, args
}
