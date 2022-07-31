package processor

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NganJason/ChatGroup-BE/internal/handler"
	"github.com/NganJason/ChatGroup-BE/internal/model"
	"github.com/NganJason/ChatGroup-BE/internal/utils"
	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
	"github.com/NganJason/ChatGroup-BE/pkg/cookies"
	"github.com/NganJason/ChatGroup-BE/vo"
)

func CreateChannelProcessor(
	ctx context.Context,
	req,
	resp interface{},
) error {
	request, ok := req.(*vo.CreateChannelRequest)
	if !ok {
		return cerr.New(
			"convert request body error",
			http.StatusBadRequest,
		)
	}

	response, ok := resp.(*vo.CreateChannelResponse)
	if !ok {
		return cerr.New(
			"convert response body error",
			http.StatusBadRequest,
		)
	}

	cookieVal := cookies.GetClientCookieValFromCtx(ctx)
	if cookieVal == nil {
		return cerr.New(
			"cookies not found",
			http.StatusForbidden,
		)
	}

	userID, err := strconv.ParseUint(*cookieVal, 10, 64)
	if err != nil {
		return cerr.New(
			fmt.Sprintf("parse cookieVal err=%s", err.Error()),
			http.StatusForbidden,
		)
	}

	p := createChannelProcessor{
		ctx:    ctx,
		userID: uint64(userID),
		req:    request,
		resp:   response,
	}

	return p.process()
}

type createChannelProcessor struct {
	ctx    context.Context
	userID uint64
	req    *vo.CreateChannelRequest
	resp   *vo.CreateChannelResponse
}

func (p *createChannelProcessor) process() error {
	err := p.validateReq()
	if err != nil {
		return cerr.New(
			err.Error(),
			http.StatusBadRequest,
		)
	}

	channelDM, err := model.NewChannelDM(p.ctx)
	if err != nil {
		return err
	}

	userChannelDM, err := model.NewUserChannelDM(p.ctx)
	if err != nil {
		return err
	}

	h := handler.NewChannelHandler(
		p.ctx,
		channelDM,
	)
	h.SetUserChannelDM(userChannelDM)

	channel, err := h.CreateChannel(
		p.req.ChannelName,
		p.req.ChannelDesc,
		utils.Uint64Ptr(p.userID),
	)
	if err != nil {
		return err
	}

	p.resp.Channel = channel

	return nil
}

func (p *createChannelProcessor) validateReq() error {
	if p.userID == 0 {
		return fmt.Errorf("userID cannot be empty")
	}

	if p.req.ChannelName == nil || *p.req.ChannelName == "" {
		return fmt.Errorf("channelName cannot be empty")
	}

	return nil
}
