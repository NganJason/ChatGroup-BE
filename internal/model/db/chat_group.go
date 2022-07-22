package db

type UserInfo struct {
	UserID         *uint64 `json:"user_id"`
	UserName       *string `json:"user_name"`
	Password       *string `json:"password"`
	HashedPassword *string `json:"hashed_password"`
	Salt           *string `json:"salt"`
	EmailAddress   *string `json:"email_address"`
	PhotoURL       *string `json:"photo_url"`
}

type UserChannel struct {
	UserID    *uint64 `json:"user_id"`
	ChannelID *uint64 `json:"channel_id"`
	Unread    *uint32 `json:"unread"`
}

type Channel struct {
	ChannelID   *uint64 `json:"channel_id"`
	ChannelName *string `json:"channel_name"`
	ChannelDesc *string `json:"channel_desc"`
	Status      *uint32 `json:"status"`
}

type Message struct {
	MessageID *uint64 `json:"message_id"`
	ChannelID *uint64 `json:"channel_id"`
	UserID    *uint64 `json:"user_id"`
	Content   *string `json:"content"`
	CreatedAt *uint64 `json:"created_at"`
}
