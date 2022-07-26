package processor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/ChatGroup-BE/internal/handler"
	"github.com/NganJason/ChatGroup-BE/internal/model"
	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
	"github.com/NganJason/ChatGroup-BE/vo"
)

func AddUsersToChannelProcessor(
	ctx context.Context,
	req,
	resp interface{},
) error {
	request, ok := resp.(*vo.AddUsersToChannelRequest)
	if !ok {
		return cerr.New(
			"convert request body error",
			http.StatusBadRequest,
		)
	}

	response, ok := resp.(*vo.AddUsersToChannelResponse)
	if !ok {
		return cerr.New(
			"convert response body error",
			http.StatusBadRequest,
		)
	}

	p := addUsersToChannelProcessor{
		ctx:  ctx,
		req:  request,
		resp: response,
	}

	return p.process()
}

type addUsersToChannelProcessor struct {
	ctx  context.Context
	req  *vo.AddUsersToChannelRequest
	resp *vo.AddUsersToChannelResponse
}

func (p *addUsersToChannelProcessor) process() error {
	err := p.validateReq()
	if err != nil {
		return cerr.New(
			err.Error(),
			http.StatusBadRequest,
		)
	}

	userChannelDM, err := model.NewUserChannelDM(p.ctx)
	if err != nil {
		return err
	}

	h := handler.NewUserChannelHandler(
		p.ctx,
		userChannelDM,
	)

	return h.AddUsersToChannel(
		p.req.ChannelID,
		p.req.UserIDs,
	)
}

func (p *addUsersToChannelProcessor) validateReq() error {
	if p.req.ChannelID == nil || *p.req.ChannelID == 0 {
		return fmt.Errorf("channelID cannot be empty")
	}

	if len(p.req.UserIDs) == 0 {
		return fmt.Errorf("userIDs and userNames cannot both be empty")
	}

	return nil
}
