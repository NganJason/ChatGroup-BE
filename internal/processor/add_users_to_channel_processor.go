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
	"github.com/NganJason/ChatGroup-BE/pkg/clog"
	"github.com/NganJason/ChatGroup-BE/pkg/cookies"
	"github.com/NganJason/ChatGroup-BE/vo"
)

func AddUsersToChannelProcessor(
	ctx context.Context,
	req,
	resp interface{},
) error {
	request, ok := req.(*vo.AddUsersToChannelRequest)
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

	cookieVal := cookies.GetClientCookieValFromCtx(ctx)
	if cookieVal == nil {
		return cerr.New(
			"cookies not found",
			http.StatusForbidden,
		)
	}

	userID, err := strconv.ParseUint(*cookieVal, 10, 64)
	if err != nil {
		clog.Error(ctx, err.Error())
		return cerr.New(
			fmt.Sprintf("parse cookieVal err=%s", err.Error()),
			http.StatusForbidden,
		)
	}

	p := addUsersToChannelProcessor{
		ctx:    ctx,
		userID: utils.Uint64Ptr(uint64(userID)),
		req:    request,
		resp:   response,
	}

	return p.process()
}

type addUsersToChannelProcessor struct {
	ctx    context.Context
	userID *uint64
	req    *vo.AddUsersToChannelRequest
	resp   *vo.AddUsersToChannelResponse
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

	userDM, err := model.NewUserDM(p.ctx)
	if err != nil {
		return err
	}

	h := handler.NewUserChannelHandler(
		p.ctx,
		userChannelDM,
	)

	h.SetUserDM(userDM)

	userIDs := make([]*uint64, 0)

	for _, id := range p.req.UserIDs {
		uint64ID := uint64(*id)
		userIDs = append(userIDs, &uint64ID)
	}

	return h.AddUsersToChannel(
		p.userID,
		p.req.ChannelID,
		userIDs,
	)
}

func (p *addUsersToChannelProcessor) validateReq() error {
	if p.userID == nil || *p.userID == 0 {
		return fmt.Errorf("userID cannot be empty")
	}

	if p.req.ChannelID == nil || *p.req.ChannelID == 0 {
		return fmt.Errorf("channelID cannot be empty")
	}

	if len(p.req.UserIDs) == 0 {
		return fmt.Errorf("userIDs and userNames cannot both be empty")
	}

	return nil
}
