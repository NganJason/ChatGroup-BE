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

func CreateMessageProcessor(
	ctx context.Context,
	req,
	resp interface{},
) error {
	request, ok := resp.(*vo.CreateMessageRequest)
	if !ok {
		return cerr.New(
			"convert request body error",
			http.StatusBadRequest,
		)
	}

	response, ok := resp.(*vo.CreateMessageResponse)
	if !ok {
		return cerr.New(
			"convert response body error",
			http.StatusBadRequest,
		)
	}

	p := createMessageProcessor{
		ctx:  ctx,
		req:  request,
		resp: response,
	}

	return p.process()
}

type createMessageProcessor struct {
	ctx  context.Context
	req  *vo.CreateMessageRequest
	resp *vo.CreateMessageResponse
}

func (p *createMessageProcessor) process() error {
	err := p.validateReq()
	if err != nil {
		return cerr.New(
			err.Error(),
			http.StatusBadRequest,
		)
	}

	userDM, err := model.NewUserDM(p.ctx)
	if err != nil {
		return err
	}

	messageDM, err := model.NewMessageDM(p.ctx)
	if err != nil {
		return err
	}

	h := handler.NewMessageHandler(
		p.ctx,
		messageDM,
		userDM,
	)

	message, err := h.CreateMessage(
		p.req.ChannelID,
		p.req.Content,
	)
	if err != nil {
		return err
	}

	p.resp.Message = message

	return nil
}

func (p *createMessageProcessor) validateReq() error {
	if p.req.ChannelID == nil || *p.req.ChannelID == 0 {
		return fmt.Errorf("channelID cannot be empty")
	}

	if p.req.Content == nil || *p.req.Content == "" {
		return fmt.Errorf("pageSize cannot be empty")
	}

	return nil
}
