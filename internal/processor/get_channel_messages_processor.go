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

	cookieVal := cookies.GetClientCookieValFromCtx(ctx)
	if cookieVal == nil {
		return cerr.New(
			"cookies not found",
			http.StatusForbidden,
		)
	}

	userID, err := strconv.Atoi(*cookieVal)
	if err != nil {
		return cerr.New(
			fmt.Sprintf("parse cookieVal err=%s", err.Error()),
			http.StatusForbidden,
		)
	}

	p := getChannelMessagesProcessor{
		ctx:    ctx,
		userID: utils.Uint64Ptr(uint64(userID)),
		req:    request,
		resp:   response,
	}

	return p.process()
}

type getChannelMessagesProcessor struct {
	ctx    context.Context
	userID *uint64
	req    *vo.GetChannelMessagesRequest
	resp   *vo.GetChannelMessagesResponse
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

	userChannelDM, err := model.NewUserChannelDM(p.ctx)
	if err != nil {
		return err
	}

	h := handler.NewMessageHandler(
		p.ctx,
		messageDM,
		userDM,
		userChannelDM,
	)

	messages, err := h.GetMessagesByChannelID(
		p.userID,
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
	if p.userID == nil || *p.userID == 0 {
		return fmt.Errorf("userID cannot be empty")
	}

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
