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

func GetChannelMembersProcessor(
	ctx context.Context,
	req,
	resp interface{},
) error {
	request, ok := req.(*vo.GetChannelMembersRequest)
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

	p := getUserChannelMembersProcessor{
		ctx:    ctx,
		userID: utils.Uint64Ptr(uint64(userID)),
		req:    request,
		resp:   response,
	}

	return p.process()
}

type getUserChannelMembersProcessor struct {
	ctx    context.Context
	userID *uint64
	req    *vo.GetChannelMembersRequest
	resp   *vo.GetChannelMembersResponse
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
		p.userID,
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
	if p.userID == nil || *p.userID == 0 {
		return fmt.Errorf("userID cannot be empty")
	}

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
