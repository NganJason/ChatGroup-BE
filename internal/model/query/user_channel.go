package query

import "strings"

type UserChannelQuery struct {
	userIDs    []*uint64
	channelIDs []*uint64
	id         *uint64
}

func NewUserChannelQuery() *UserChannelQuery {
	return &UserChannelQuery{}
}

func (q *UserChannelQuery) ID(
	id *uint64,
) *UserChannelQuery {
	q.id = id

	return q
}

func (q *UserChannelQuery) ChannelID(
	channelID *uint64,
) *UserChannelQuery {
	q.channelIDs = append(q.channelIDs, channelID)

	return q
}

func (q *UserChannelQuery) ChannelIDs(
	ChannelIDs []*uint64,
) *UserChannelQuery {
	for _, id := range ChannelIDs {
		q.channelIDs = append(q.channelIDs, id)
	}

	return q
}

func (q *UserChannelQuery) UserID(
	userID *uint64,
) *UserChannelQuery {
	q.userIDs = append(q.userIDs, userID)

	return q
}

func (q *UserChannelQuery) UserIDs(
	userIDs []*uint64,
) *UserChannelQuery {
	for _, userID := range userIDs {
		q.userIDs = append(q.userIDs, userID)
	}

	return q
}

func (q *UserChannelQuery) Build() (
	wheres string,
	args []interface{},
) {
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
			args = append(args, *id)
		}
	}

	if len(q.userIDs) != 0 {
		inCondition := "user_id IN (?"

		for i := 1; i < len(q.userIDs); i++ {
			inCondition = inCondition + ",?"
		}
		inCondition = inCondition + ")"
		whereCols = append(whereCols, inCondition)

		for _, id := range q.userIDs {
			args = append(args, *id)
		}
	}

	wheres = strings.Join(whereCols, " AND ")

	return wheres, args
}
