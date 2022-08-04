package processor

import (
	"context"
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

func ValidateAuthProcessor(
	ctx context.Context,
	req,
	resp interface{},
) error {
	response, ok := resp.(*vo.ValidateAuthResponse)
	if !ok {
		return cerr.New(
			"convert response body error",
			http.StatusBadRequest,
		)
	}

	cookieVal := cookies.GetClientCookieValFromCtx(ctx)
	if cookieVal == nil {
		clog.Info(ctx, "cookie is nil")
		return nil
	}

	userID, err := strconv.ParseUint(*cookieVal, 10, 64)
	if err != nil {
		clog.Error(ctx, err.Error())
		return nil
	}

	p := validateAuthProcessor{
		ctx: ctx,
		userID: utils.Uint64Ptr(
			uint64(userID)),
		resp: response,
	}

	return p.process()
}

type validateAuthProcessor struct {
	ctx    context.Context
	userID *uint64
	resp   *vo.ValidateAuthResponse
}

func (p *validateAuthProcessor) process() error {
	if p.userID == nil || *p.userID == 0 {
		clog.Info(p.ctx, "userID is empty")
		return nil
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

	p.resp.User = user

	return nil
}
