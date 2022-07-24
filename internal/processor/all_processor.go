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
			Req:       vo.GetUserInfoRequest{},
			Resp:      vo.GetUserInfoResponse{},
		},
		{
			Path:      "/api/user/channels",
			Processor: GetUserChannelsProcessor,
			Req:       vo.GetUserChannelsRequest{},
			Resp:      vo.GetUserChannelsResponse{},
		},
		{
			Path:      "/api/channel/messages",
			Processor: GetChannelMessagesProcessor,
			Req:       vo.GetChannelMessagesRequest{},
			Resp:      vo.GetChannelMessagesResponse{},
		},
		{
			Path:      "/api/channel/members",
			Processor: GetChannelMembersProcessor,
			Req:       vo.GetChannelMembersRequest{},
			Resp:      vo.GetChannelMembersResponse{},
		},
		{
			Path: "/api/message/create",
			Req:  vo.CreateMessageRequest{},
			Resp: vo.CreateMessageResponse{},
		},
	}
}
