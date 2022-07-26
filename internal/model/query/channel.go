package query

import "strings"

type ChannelQuery struct {
	channelIDs   []*uint64
	channelNames []*string
	id           *uint64
}

func NewChannelQuery() *ChannelQuery {
	return &ChannelQuery{}
}

func (q *ChannelQuery) ChannelID(
	channelID *uint64,
) *ChannelQuery {
	q.channelIDs = append(q.channelIDs, channelID)

	return q
}

func (q *ChannelQuery) ChannelIDs(
	channelIDs []*uint64,
) *ChannelQuery {
	for _, id := range channelIDs {
		q.channelIDs = append(q.channelIDs, id)
	}

	return q
}

func (q *ChannelQuery) ChannelName(
	channelName *string,
) *ChannelQuery {
	q.channelNames = append(q.channelNames, channelName)

	return q
}

func (q *ChannelQuery) ChannelNames(
	channelNames []*string,
) *ChannelQuery {
	for _, channelName := range channelNames {
		q.channelNames = append(q.channelNames, channelName)
	}

	return q
}

func (q *ChannelQuery) ID(
	id *uint64,
) *ChannelQuery {
	q.id = id

	return q
}

func (q *ChannelQuery) Build() (wheres string, args []interface{}) {
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

	if len(q.channelNames) != 0 {
		inCondition := "channel_name IN (?"

		for i := 1; i < len(q.channelNames); i++ {
			inCondition = inCondition + ",?"
		}
		inCondition = inCondition + ")"
		whereCols = append(whereCols, inCondition)

		for _, channelName := range q.channelNames {
			args = append(args, channelName)
		}
	}

	wheres = strings.Join(whereCols, " AND ")

	return wheres, args
}
