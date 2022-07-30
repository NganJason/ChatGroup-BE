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

func CreateMessageProcessor(
	ctx context.Context,
	req,
	resp interface{},
) error {
	request, ok := req.(*vo.CreateMessageRequest)
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

	p := createMessageProcessor{
		ctx:    ctx,
		userID: utils.Uint64Ptr(uint64(userID)),
		req:    request,
		resp:   response,
	}

	return p.process()
}

type createMessageProcessor struct {
	ctx    context.Context
	userID *uint64
	req    *vo.CreateMessageRequest
	resp   *vo.CreateMessageResponse
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

	message, err := h.CreateMessage(
		p.req.ChannelID,
		p.req.Content,
		p.userID,
	)
	if err != nil {
		return err
	}

	p.resp.Message = message

	return nil
}

func (p *createMessageProcessor) validateReq() error {
	if p.userID == nil || *p.userID == 0 {
		return fmt.Errorf("userID cannot be empty")
	}

	if p.req.ChannelID == nil || *p.req.ChannelID == 0 {
		return fmt.Errorf("channelID cannot be empty")
	}

	if p.req.Content == nil || *p.req.Content == "" {
		return fmt.Errorf("pageSize cannot be empty")
	}

	return nil
}
