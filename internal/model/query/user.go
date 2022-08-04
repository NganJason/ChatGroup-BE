package query

import (
	"fmt"
	"strings"
)

type UserQuery struct {
	userIDs   []*uint64
	userNames []*string
	id        *uint64
	keyword   *string
}

func NewUserQuery() *UserQuery {
	return &UserQuery{}
}

func (q *UserQuery) ID(
	id *uint64,
) *UserQuery {
	q.id = id

	return q
}

func (q *UserQuery) UserID(
	userID *uint64,
) *UserQuery {
	q.userIDs = append(q.userIDs, userID)

	return q
}

func (q *UserQuery) UserName(
	userName *string,
) *UserQuery {
	q.userNames = append(q.userNames, userName)

	return q
}

func (q *UserQuery) UserIDs(
	userIDs []*uint64,
) *UserQuery {
	for _, id := range userIDs {
		q.userIDs = append(q.userIDs, id)
	}

	return q
}

func (q *UserQuery) Keyword(
	keyword *string,
) *UserQuery {
	q.keyword = keyword

	return q
}

func (q *UserQuery) Build() (wheres string, args []interface{}) {
	whereCols := make([]string, 0)

	if q.id != nil {
		whereCols = append(whereCols, "id = ?")
		args = append(args, *q.id)
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

	if len(q.userNames) != 0 {
		inCondition := "username IN (?"

		for i := 1; i < len(q.userNames); i++ {
			inCondition = inCondition + ",?"
		}
		inCondition = inCondition + ")"
		whereCols = append(whereCols, inCondition)

		for _, userName := range q.userNames {
			args = append(args, *userName)
		}
	}

	if q.keyword != nil {
		whereCols = append(whereCols,
			fmt.Sprintf("username LIKE '%s%%'", *q.keyword),
		)
	}

	wheres = strings.Join(whereCols, " AND ")

	return wheres, args
}
