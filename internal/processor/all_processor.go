package processor

import (
	"context"
	"github.com/NganJason/ChatGroup-BE/vo"
)

type ProcessorConfig struct {
	Path      string
	Processor func(ctx context.Context, req, resp interface{}) error
	Req       interface{}
	Resp      interface{}
	NeedAuth  bool
}

func GetAllProcessors() []ProcessorConfig {
	return []ProcessorConfig{
		{
			Path:      "/api/healthcheck",
			Processor: HealthCheckProcessor,
			Req:       &vo.HealthCheckRequest{},
			Resp:      &vo.HealthCheckResponse{},
		},
		{
			Path:      "/api/auth/login",
			Processor: LoginProcessor,
			Req:       &vo.AuthLoginRequest{},
			Resp:      &vo.AuthLoginResponse{},
		},
		{
			Path:      "/api/auth/signup",
			Processor: SignupProcessor,
			Req:       &vo.AuthSignupRequest{},
			Resp:      &vo.AuthSignupResponse{},
		},
		{
			Path:      "/api/user/info",
			Processor: GetUserInfoProcessor,
			Req:       &vo.GetUserInfoRequest{},
			Resp:      &vo.GetUserInfoResponse{},
			NeedAuth: true,
		},
		{
			Path:      "/api/user/channels",
			Processor: GetUserChannelsProcessor,
			Req:       &vo.GetUserChannelsRequest{},
			Resp:      &vo.GetUserChannelsResponse{},
		},
		{
			Path:      "/api/channel/create",
			Processor: CreateChannelProcessor,
			Req:       &vo.CreateChannelRequest{},
			Resp:      &vo.CreateChannelResponse{},
		},
		{
			Path:      "/api/channel/messages",
			Processor: GetChannelMessagesProcessor,
			Req:       &vo.GetChannelMessagesRequest{},
			Resp:      &vo.GetChannelMessagesResponse{},
		},
		{
			Path:      "/api/channel/members",
			Processor: GetChannelMembersProcessor,
			Req:       &vo.GetChannelMembersRequest{},
			Resp:      &vo.GetChannelMembersResponse{},
		},
		{
			Path:      "/api/channel/add_users",
			Processor: AddUsersToChannelProcessor,
			Req:       &vo.AddUsersToChannelRequest{},
			Resp:      &vo.AddUsersToChannelResponse{},
		},
		{
			Path:      "/api/message/create",
			Processor: CreateMessageProcessor,
			Req:       &vo.CreateMessageRequest{},
			Resp:      &vo.CreateMessageResponse{},
		},
	}
}
