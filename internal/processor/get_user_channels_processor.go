package processor

import (
	"context"
	"fmt"
	"github.com/NganJason/ChatGroup-BE/internal/handler"
	"github.com/NganJason/ChatGroup-BE/internal/model"
	"github.com/NganJason/ChatGroup-BE/internal/utils"
	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
	"github.com/NganJason/ChatGroup-BE/pkg/cookies"
	"github.com/NganJason/ChatGroup-BE/vo"
	"net/http"
	"strconv"
)

func GetUserChannelsProcessor(ctx context.Context, req, resp interface{}) error {
	response, ok := resp.(*vo.GetUserChannelsResponse)
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

	p := getUserChannelsProcessor{
		ctx: ctx,
		userID: utils.Uint64Ptr(
			uint64(userID)),
		resp: response,
	}

	return p.process()
}

type getUserChannelsProcessor struct {
	ctx    context.Context
	userID *uint64
	resp   *vo.GetUserChannelsResponse
}

func (p *getUserChannelsProcessor) process() error {
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

	channels, err := h.GetUserChannels(p.userID)
	if err != nil {
		return err
	}

	p.resp.Channels = channels

	return nil
}

func (p *getUserChannelsProcessor) validateReq() error {
	if p.userID == nil {
		return fmt.Errorf(
			"userID cannot be empty",
		)
	}

	return nil
}
