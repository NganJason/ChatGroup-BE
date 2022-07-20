package processor

import (
	"context"
	"github.com/NganJason/ChatGroup-BE/internal/vo"
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
	}
}
