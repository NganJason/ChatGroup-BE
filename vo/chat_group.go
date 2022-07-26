package vo

type HealthCheckRequest struct{}

type HealthCheckResponse struct {
	DebugMsg *string `json:"debug_msg"`
	Message  *string `json:"message"`
}

type AuthLoginRequest struct {
	UserName *string `json:"user_name"`
	Password *string `json:"password"`
}

type AuthLoginResponse struct {
	DebugMsg *string `json:"debug_msg"`
	User     *User   `json:"user_info"`
}

type AuthSignupRequest struct {
	UserName *string `json:"user_name"`
	Password *string `json:"password"`
}

type AuthSignupResponse struct {
	DebugMsg *string `json:"debug_msg"`
	User     *User   `json:"user_info"`
}

type GetUserInfoRequest struct{}

type GetUserInfoResponse struct {
	DebugMsg *string `json:"debug_msg"`
	UserInfo *User   `json:"user_info"`
}

type GetUserChannelsRequest struct{}

type GetUserChannelsResponse struct {
	DebugMsg *string   `json:"debug_msg"`
	Channels []Channel `json:"channels"`
}

type CreateChannelRequest struct {
	UserID      *uint64 `json:"user_id"`
	ChannelName *string `json:"channel_name"`
	ChannelDesc *string `json:"channel_desc"`
}

type CreateChannelResponse struct {
	DebugMsg *string  `json:"debug_msg"`
	Channel  *Channel `json:"channel"`
}

type GetChannelMessagesRequest struct {
	ChannelID    *uint64 `json:"channel_id"`
	FromUnixTime *uint64 `json:"from_unix_time"`
	ToUnixTime   *uint64 `json:"to_unix_time"`
}

type GetChannelMessagesResponse struct {
	DebugMsg *string   `json:"debug_msg"`
	Messages []Message `json:"messages"`
}

type GetChannelMembersRequest struct {
	ChannelID  *uint64 `json:"channel_id"`
	PageSize   *uint32 `json:"page_size"`
	PageNumber *uint32 `json:"page_number"`
}

type GetChannelMembersResponse struct {
	DebugMsg *string `json:"debug_msg"`
	Members  []User  `json:"members"`
}

type CreateMessageRequest struct {
	ChannelID *uint64 `json:"channel_id"`
	Content   *string `json:"content"`
}

type CreateMessageResponse struct {
	DebugMsg *string  `json:"debug_msg"`
	Message  *Message `json:"message"`
}

type AddUsersToChannelRequest struct {
	ChannelID *uint64   `json:"channel_id"`
	UserIDs   []*uint64 `json:"user_ids"`
}

type AddUsersToChannelResponse struct {
	DebugMsg *string `json:"debug_msg"`
}

type User struct {
	UserID       *uint64 `json:"user_id"`
	UserName     *string `json:"user_name"`
	EmailAddress *string `json:"email_address"`
	PhotoURL     *string `json:"photo_url"`
}

type Channel struct {
	ChannelID   *uint64 `json:"channel_id"`
	ChannelName *string `json:"channel_name"`
	ChannelDesc *string `json:"channel_desc"`
	Unread      *uint32 `json:"unread"`
}

type Message struct {
	MessageID *uint64 `json:"message_id"`
	ChannelID *uint64 `json:"channel_id"`
	Content   *string `json:"content"`
	CreatedAt *uint64 `json:"created_at"`
	Sender    *User   `json:"sender"`
}
