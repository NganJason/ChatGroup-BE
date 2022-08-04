package processor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/ChatGroup-BE/internal/handler"
	"github.com/NganJason/ChatGroup-BE/internal/model"
	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
	"github.com/NganJason/ChatGroup-BE/vo"
)

func SearchUsersProcessor(
	ctx context.Context,
	req,
	resp interface{},
) error {
	request, ok := req.(*vo.SearchUsersRequest)
	if !ok {
		return cerr.New(
			"convert request body error",
			http.StatusBadRequest,
		)
	}

	response, ok := resp.(*vo.SearchUsersResponse)
	if !ok {
		return cerr.New(
			"convert response body error",
			http.StatusBadRequest,
		)
	}

	p := searchUsersProcessor{
		ctx:  ctx,
		req:  request,
		resp: response,
	}

	return p.process()
}

type searchUsersProcessor struct {
	ctx  context.Context
	req  *vo.SearchUsersRequest
	resp *vo.SearchUsersResponse
}

func (p *searchUsersProcessor) process() error {
	if err := p.validateReq(); err != nil {
		return err
	}

	userDM, err := model.NewUserDM(p.ctx)
	if err != nil {
		return err
	}

	h := handler.NewUserHandler(
		p.ctx,
		userDM,
	)

	users, err := h.SearchUsers(
		p.req.Keyword,
	)
	if err != nil {
		return err
	}

	p.resp.Users = users

	return nil
}

func (p *searchUsersProcessor) validateReq() error {
	if p.req.Keyword == nil || *p.req.Keyword == "" {
		return fmt.Errorf("keyword cannot be empty")
	}

	return nil
}
