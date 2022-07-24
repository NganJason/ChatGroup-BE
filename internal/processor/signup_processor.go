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
	"github.com/NganJason/ChatGroup-BE/vo"
)

func SignupProcessor(ctx context.Context, req, resp interface{}) error {
	request, ok := req.(*vo.AuthSignupRequest)
	if !ok {
		return cerr.New(
			"convert request body error",
			http.StatusBadRequest,
		)
	}

	response, ok := resp.(*vo.AuthSignupResponse)
	if !ok {
		return cerr.New(
			"convert response body error",
			http.StatusBadRequest,
		)
	}

	p := signupProcessor{
		ctx:  ctx,
		req:  request,
		resp: response,
	}

	err := p.process()
	if err != nil {
		return err
	}

	err = utils.GenerateTokenAndAddCookies(
		p.ctx,
		strconv.Itoa(int(*p.resp.User.UserID)),
	)
	if err != nil {
		return cerr.New(
			err.Error(),
			http.StatusBadGateway,
		)
	}

	return nil
}

type signupProcessor struct {
	ctx  context.Context
	req  *vo.AuthSignupRequest
	resp *vo.AuthSignupResponse
}

func (p *signupProcessor) process() error {
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

	h := handler.NewAuthHandler(
		p.ctx,
		userDM,
	)

	user, err := h.Signup(
		p.req.UserName,
		p.req.Password,
	)
	if err != nil {
		return err
	}

	p.resp.User = user

	return nil
}

func (p *signupProcessor) validateReq() error {
	if p.req.UserName == nil || p.req.Password == nil {
		return fmt.Errorf("username or password cannot be empty")
	}

	if *p.req.UserName == "" || *p.req.Password == "" {
		return fmt.Errorf("username or password cannot be empty")
	}

	return nil
}
