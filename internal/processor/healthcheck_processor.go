package processor

import (
	"context"
	"net/http"

	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
	"github.com/NganJason/ChatGroup-BE/vo"
	"google.golang.org/protobuf/proto"
)

func HealthCheckProcessor(ctx context.Context, req, resp interface{}) error {
	_, ok := req.(*vo.HealthCheckRequest)
	if !ok {
		return cerr.New(
			"convert request body error",
			http.StatusBadRequest,
		)
	}

	response, ok := resp.(*vo.HealthCheckResponse)
	if !ok {
		return cerr.New(
			"convert response body error",
			http.StatusBadRequest,
		)
	}

	response.Message = proto.String("I am healthy")

	return nil
}
