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

func GetChannelMessagesProcessor(
	ctx context.Context,
	req,
	resp interface{},
) error {
	request, ok := req.(*vo.GetChannelMessagesRequest)
	if !ok {
		return cerr.New(
			"convert request body error",
			http.StatusBadRequest,
		)
	}

	response, ok := resp.(*vo.GetChannelMessagesResponse)
	if !ok {
		return cerr.New(
			"convert response body error",
			http.StatusBadRequest,
		)
	}

	p := getChannelMessagesProcessor{
		ctx:  ctx,
		req:  request,
		resp: response,
	}

	return p.process()
}

type getChannelMessagesProcessor struct {
	ctx  context.Context
	req  *vo.GetChannelMessagesRequest
	resp *vo.GetChannelMessagesResponse
}

func (p *getChannelMessagesProcessor) process() error {
	err := p.validateReq()
	if err != nil {
		return cerr.New(
			err.Error(),
			http.StatusBadRequest,
		)
	}

	messageDM, err := model.NewMessageDM(p.ctx)
	if err != nil {
		return err
	}

	userDM, err := model.NewUserDM(p.ctx)
	if err != nil {
		return err
	}

	h := handler.NewMessageHandler(
		p.ctx,
		messageDM,
		userDM,
	)

	messages, err := h.GetMessagesByChannelID(
		p.req.ChannelID,
		p.req.FromUnixTime,
		p.req.ToUnixTime,
	)
	if err != nil {
		return err
	}

	p.resp.Messages = messages

	return nil
}

func (p *getChannelMessagesProcessor) validateReq() error {
	if p.req.ChannelID == nil || *p.req.ChannelID == 0 {
		return fmt.Errorf("channelID cannot be empty")
	}

	if p.req.FromUnixTime == nil || *p.req.FromUnixTime == 0 {
		return fmt.Errorf("fromUnixTime cannot be empty")
	}

	if p.req.ToUnixTime == nil || *p.req.ToUnixTime == 0 {
		return fmt.Errorf("toUnixTime cannot be empty")
	}

	return nil
}
