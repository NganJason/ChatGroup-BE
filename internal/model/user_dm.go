package model

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/NganJason/ChatGroup-BE/internal/model/db"
	"github.com/NganJason/ChatGroup-BE/internal/model/query"
	"github.com/NganJason/ChatGroup-BE/internal/utils"
	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
)

type UserDM struct {
	ctx context.Context
	db  *sql.DB
}

func NewUserDM(ctx context.Context) (User, error) {
	return &UserDM{
		ctx: ctx,
	}, nil
}

func (dm *UserDM) GetUser(
	userID *uint64,
	userName *string,
	id *uint64,
) (user *db.User, err error) {
	query := query.NewUserQuery()

	if userID != nil {
		query.UserID(userID)
	}

	if userName != nil {
		query.UserName(userName)
	}

	if id != nil {
		query.ID(id)
	}

	baseQuery := fmt.Sprintf(
		`SELECT * FROM %s WHERE `,
		dm.getTableName(),
	)

	wheres, args := query.Build()

	err = dm.db.QueryRow(
		baseQuery+wheres,
		args...,
	).Scan(
		&user.ID,
		&user.UserID,
		&user.UserName,
		&user.EmailAddress,
		&user.HashedPassword,
		&user.PhotoURL,
		&user.Salt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, cerr.New(
			fmt.Sprintf("query user err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	return user, nil
}

func (dm *UserDM) GetUsers(userIDs []*uint64) (users []*db.User, err error) {
	if len(userIDs) == 0 {
		return nil, cerr.New(
			"userIDs cannot be empty",
			http.StatusBadRequest,
		)
	}

	baseQuery := fmt.Sprintf(
		`SELECT * from %s WHERE `,
		dm.getTableName(),
	)

	query := query.NewUserQuery().UserIDs(userIDs)
	wheres, args := query.Build()

	rows, err := dm.db.Query(
		baseQuery+wheres,
		args...,
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("query users from db err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	for rows.Next() {
		var user *db.User

		if err := rows.Scan(
			&user.ID,
			&user.UserID,
			&user.UserName,
			&user.EmailAddress,
			&user.HashedPassword,
			&user.Salt,
			&user.PhotoURL,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			if err == sql.ErrNoRows {
				return users, nil
			}

			return nil, cerr.New(
				fmt.Sprintf("query users from db err=%s", err.Error()),
				http.StatusBadGateway,
			)
		}

		users = append(users, user)
	}

	return users, nil
}

func (dm *UserDM) CreateUser(req *CreateUserReq) (user *db.User, err error) {
	query := fmt.Sprintf(
		`
		INSERT INTO %s 
		(user_id, user_name, hashed_password, salt, email_address, photo_url) 
		VALUES(?, ?, ?, ?, ?, ?)
		`, dm.getTableName(),
	)

	result, err := dm.db.Exec(
		query,
		req.UserID,
		req.UserName,
		req.HashedPassword,
		req.EmailAddress,
		req.PhotoURL,
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("insert user into db err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	lastInsertID, _ := result.LastInsertId()

	user, err = dm.GetUser(
		nil,
		nil,
		utils.Uint64Ptr(uint64(lastInsertID)),
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("refetch user from db err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	return user, nil
}

func (dm *UserDM) UpdateUser(req *UpdateUserReq) (user *db.User, err error) {
	if req.UserID == 0 {
		return nil, cerr.New(
			"userID cannot be empty for update",
			http.StatusBadRequest,
		)
	}

	tx, err := dm.db.BeginTx(dm.ctx, nil)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("begin tx for update err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}
	defer tx.Rollback()

	baseQuery := fmt.Sprintf(
		`SELECT * from %s WHERE `,
		dm.getTableName(),
	)

	query := query.NewUserQuery().UserID(utils.Uint64Ptr(req.UserID))
	wheres, args := query.Build()
	finalQuery := baseQuery + wheres + "FOR UPDATE"

	var existingUser *db.User
	err = tx.QueryRowContext(
		dm.ctx,
		finalQuery,
		args...,
	).Scan(
		&existingUser.ID,
		&existingUser.UserID,
		&existingUser.UserName,
		&existingUser.EmailAddress,
		&existingUser.HashedPassword,
		&existingUser.Salt,
		&existingUser.PhotoURL,
		&existingUser.CreatedAt,
		&existingUser.UpdatedAt,
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("get existing user err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	if existingUser == nil {
		return nil, cerr.New(
			"user does not exist for update",
			http.StatusBadRequest,
		)
	}

	if req.UserName != nil {
		existingUser.UserName = req.UserName
	}

	if req.HashedPassword != nil {
		existingUser.HashedPassword = req.HashedPassword
	}

	if req.Salt != nil {
		existingUser.Salt = req.Salt
	}

	if req.EmailAddress != nil {
		existingUser.EmailAddress = req.EmailAddress
	}

	if req.PhotoURL != nil {
		existingUser.PhotoURL = req.PhotoURL
	}

	existingUser.UpdatedAt = time.Now()

	updateQuery := fmt.Sprintf(
		`
		UPDATE %S
		SET user_name = ?, hashed_password = ?, salt = ?, email_address = ?, photo_url = ?
		WHERE user_id = ?,
		`,
		dm.getTableName(),
	)

	_, err = tx.ExecContext(
		dm.ctx,
		updateQuery,
		existingUser.UserName,
		existingUser.HashedPassword,
		existingUser.Salt,
		existingUser.EmailAddress,
		existingUser.PhotoURL,
		existingUser.UserID,
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("update user err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	err = tx.Commit()
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("commit transaction err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	return existingUser, nil
}

func (dm *UserDM) getTableName() string {
	return "user_tab"
}
