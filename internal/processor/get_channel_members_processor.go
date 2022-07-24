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

func GetChannelMembersProcessor(
	ctx context.Context,
	req,
	resp interface{},
) error {
	request, ok := resp.(*vo.GetChannelMembersRequest)
	if !ok {
		return cerr.New(
			"convert request body error",
			http.StatusBadRequest,
		)
	}

	response, ok := resp.(*vo.GetChannelMembersResponse)
	if !ok {
		return cerr.New(
			"convert response body error",
			http.StatusBadRequest,
		)
	}

	p := getUserChannelMembersProcessor{
		ctx:  ctx,
		req:  request,
		resp: response,
	}

	return p.process()
}

type getUserChannelMembersProcessor struct {
	ctx  context.Context
	req  *vo.GetChannelMembersRequest
	resp *vo.GetChannelMembersResponse
}

func (p *getUserChannelMembersProcessor) process() error {
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

	userDM, err := model.NewUserDM(p.ctx)
	if err != nil {
		return err
	}

	h := handler.NewUserChannelHandler(
		p.ctx,
		userChannelDM,
	)
	h.SetUserDM(userDM)

	users, err := h.GetChannelUsers(
		p.req.ChannelID,
		p.req.PageSize,
		p.req.PageNumber,
	)
	if err != nil {
		return err
	}

	p.resp.Members = users

	return nil
}

func (p *getUserChannelMembersProcessor) validateReq() error {
	if p.req.ChannelID == nil || *p.req.ChannelID == 0 {
		return fmt.Errorf("channelID cannot be empty")
	}

	if p.req.PageSize == nil || *p.req.PageSize == 0 {
		return fmt.Errorf("pageSize cannot be empty")
	}

	if p.req.PageNumber == nil || *p.req.PageNumber == 0 {
		return fmt.Errorf("pageNumber cannot be empty")
	}

	return nil
}
