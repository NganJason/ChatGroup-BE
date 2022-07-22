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

func GetUserInfoProcessor(ctx context.Context, req, resp interface{}) error {
	response, ok := resp.(*vo.GetUserInfoResponse)
	if !ok {
		return cerr.New(
			"convert response body error",
			http.StatusBadRequest,
		)
	}

	p := getUserInfoProcessor{
		ctx:  ctx,
		resp: response,
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

	p.userID = utils.Uint64Ptr(
		uint64(userID),
	)

	return p.process()
}

type getUserInfoProcessor struct {
	ctx    context.Context
	userID *uint64
	resp   *vo.GetUserInfoResponse
}

func (p *getUserInfoProcessor) process() error {
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

	h := handler.NewUserHandler(p.ctx, userDM)
	user, err := h.GetUser(p.userID)
	if err != nil {
		return err
	}

	p.resp.UserInfo = user

	return nil
}

func (p *getUserInfoProcessor) validateReq() error {
	if p.userID == nil {
		return fmt.Errorf(
			"userID cannot be empty",
		)
	}

	return nil
}
