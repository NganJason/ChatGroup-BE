package cerr

import (
	"context"
	"errors"
	"net/http"
)

type cerr struct {
	error
	code int
}

type ctxKey string

const (
	errKey = ctxKey("error")
)

func New(msg string, code int) error {
	return &cerr{
		error: errors.New(msg),
		code:  code,
	}
}

func Code(err error) int {
	if err == nil {
		return http.StatusOK
	}

	if cerr, ok := err.(*cerr); ok {
		return cerr.code
	} else {
		return http.StatusBadGateway
	}
}

func AddErrToCtx(ctx context.Context, err error) context.Context {
	return context.WithValue(ctx, errKey, err)
}

func GetErrFromCtx(ctx context.Context) error {
	val := ctx.Value(errKey)
	if val == nil {
		return nil
	}

	return val.(*cerr)
}
