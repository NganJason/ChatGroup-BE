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

func UploadImageProcessor(
	ctx context.Context,
	req,
	resp interface{},
) error {
	response, ok := resp.(*vo.UploadImageResponse)
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

	p := uploadImageProcessor{
		ctx:    ctx,
		userID: uint64(userID),
		resp:   response,
	}

	return p.process()
}

type uploadImageProcessor struct {
	ctx    context.Context
	userID uint64
	resp   *vo.UploadImageResponse
}

func (p *uploadImageProcessor) process() error {
	userDM, err := model.NewUserDM(p.ctx)
	if err != nil {
		return err
	}

	file := p.ctx.Value(utils.ImageCtxKey)
	if file == nil {
		return cerr.New(
			"cannot parse file",
			http.StatusBadGateway,
		)
	}

	fileBytes, _ := file.([]byte)
	if fileBytes == nil {
		return cerr.New(
			"assert file error",
			http.StatusBadGateway,
		)
	}

	h := handler.NewUserHandler(
		p.ctx,
		userDM,
	)

	url, err := h.UploadImage(
		fileBytes,
		p.userID,
	)
	if err != nil {
		return err
	}

	p.resp.Url = utils.StrPtr(url)

	return nil
}
