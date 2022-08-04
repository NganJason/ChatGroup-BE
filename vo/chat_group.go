package vo

import (
	"encoding/json"
	"strconv"
)

type HealthCheckRequest struct{}

type HealthCheckResponse struct {
	DebugMsg *string `json:"debug_msg"`
	Message  *string `json:"message"`
}

type ValidateAuthRequest struct{}

type ValidateAuthResponse struct {
	DebugMsg *string `json:"debug_msg"`
	User     *User   `json:"user_info"`
}
type AuthLoginRequest struct {
	UserName *string `json:"username"`
	Password *string `json:"password"`
}

type AuthLoginResponse struct {
	DebugMsg *string `json:"debug_msg"`
	User     *User   `json:"user_info"`
}

type AuthSignupRequest struct {
	UserName *string `json:"username"`
	Password *string `json:"password"`
}

type AuthSignupResponse struct {
	DebugMsg *string `json:"debug_msg"`
	User     *User   `json:"user_info"`
}

type AuthLogoutRequest struct{}

type AuthLogoutResponse struct {
	DebugMsg *string `json:"debug_msg"`
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

type SearchUsersRequest struct {
	Keyword *string `json:"keyword"`
}

type SearchUsersResponse struct {
	DebugMsg *string `json:"debug_msg"`
	Users    []User  `json:"users"`
}

type CreateChannelRequest struct {
	ChannelName *string `json:"channel_name"`
	ChannelDesc *string `json:"channel_desc"`
}

type CreateChannelResponse struct {
	DebugMsg *string  `json:"debug_msg"`
	Channel  *Channel `json:"channel"`
}

type GetChannelMessagesRequest struct {
	ChannelID    *uint64 `json:"channel_id,string"`
	FromUnixTime *uint64 `json:"from_unix_time"`
	PageSize   *uint32 `json:"page_size"`
}

type GetChannelMessagesResponse struct {
	DebugMsg *string   `json:"debug_msg"`
	Messages []Message `json:"messages"`
}

type GetChannelMembersRequest struct {
	ChannelID  *uint64 `json:"channel_id,string"`
	PageSize   *uint32 `json:"page_size"`
	PageNumber *uint32 `json:"page_number"`
}

type GetChannelMembersResponse struct {
	DebugMsg *string `json:"debug_msg"`
	Members  []User  `json:"members"`
}

type CreateMessageRequest struct {
	ChannelID *uint64 `json:"channel_id,string"`
	Content   *string `json:"content"`
}

type CreateMessageResponse struct {
	DebugMsg *string  `json:"debug_msg"`
	Message  *Message `json:"message"`
}

type AddUsersToChannelRequest struct {
	ChannelID *uint64      `json:"channel_id,string"`
	UserIDs   []*Uint64Str `json:"user_ids,string"`
}

type AddUsersToChannelResponse struct {
	DebugMsg *string `json:"debug_msg"`
}

type User struct {
	UserID       *uint64 `json:"user_id,string"`
	UserName     *string `json:"username"`
	EmailAddress *string `json:"email_address"`
	PhotoURL     *string `json:"photo_url"`
}

type Channel struct {
	ChannelID   *uint64 `json:"channel_id,string"`
	ChannelName *string `json:"channel_name"`
	ChannelDesc *string `json:"channel_desc"`
	Unread      *uint32 `json:"unread"`
}

type Message struct {
	MessageID *uint64 `json:"message_id"`
	ChannelID *uint64 `json:"channel_id,string"`
	Content   *string `json:"content"`
	CreatedAt *uint64 `json:"created_at"`
	Sender    *User   `json:"sender"`
}

type Uint64Str uint64

func (i Uint64Str) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatInt(int64(i), 10))
}

func (i *Uint64Str) UnmarshalJSON(b []byte) error {
	// Try string first
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		value, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return err
		}
		*i = Uint64Str(value)
		return nil
	}

	// Fallback to number
	return json.Unmarshal(b, (*uint64)(i))
}
