package db

import "time"

type User struct {
	ID             *uint64 `json:"id"`
	UserID         *uint64 `json:"user_id"`
	UserName       *string `json:"user_name"`
	HashedPassword *string `json:"hashed_password"`
	Salt           *string `json:"salt"`
	EmailAddress   *string `json:"email_address"`
	PhotoURL       *string `json:"photo_url"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserChannel struct {
	ID        *uint64 `json:"id"`
	UserID    *uint64 `json:"user_id"`
	ChannelID *uint64 `json:"channel_id"`
	Unread    *uint32 `json:"unread"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Channel struct {
	ID          *uint64 `json:"id"`
	ChannelID   *uint64 `json:"channel_id"`
	ChannelName *string `json:"channel_name"`
	ChannelDesc *string `json:"channel_desc"`
	Status      *uint32 `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Message struct {
	ID        *uint64 `json:"id"`
	MessageID *uint64 `json:"message_id"`
	ChannelID *uint64 `json:"channel_id"`
	UserID    *uint64 `json:"user_id"`
	Content   *string `json:"content"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
