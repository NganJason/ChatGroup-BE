package cookies

import (
	"context"
	"net/http"
)

type cookieOption func(*http.Cookie)

type cookieKey string

var CKey cookieKey

const (
	ClientCookieCtxKey = cookieKey("cookie_from_client")
	ServerCookieCtxKey = cookieKey("cookie_from_server")

	ClientCookieValCtxKey = cookieKey("cookie_val_from_client")
)

func SetCookieKey(key string) {
	CKey = cookieKey(key)
}

func GetCookieKey() string {
	if CKey == "" {
		return "default_cookie_key"
	}

	return string(CKey)
}

func AddClientCookieValToCtx(ctx context.Context, val *string) context.Context {
	return context.WithValue(ctx, ClientCookieValCtxKey, val)
}

func GetClientCookieValFromCtx(ctx context.Context) *string {
	c := ctx.Value(ClientCookieValCtxKey)
	if c == nil {
		return nil
	}

	return c.(*string)
}

func InitServerCookie(ctx context.Context) context.Context {
	serverCookie := new(http.Cookie)
	ctx = context.WithValue(ctx, ServerCookieCtxKey, serverCookie)

	return ctx
}

func AddServerCookieToCtx(ctx context.Context, cookie *http.Cookie) {
	c := ctx.Value(ServerCookieCtxKey)
	existingCookie := c.(*http.Cookie)

	*existingCookie = *cookie
}

func GetServerCookieFromCtx(ctx context.Context) *http.Cookie {
	c := ctx.Value(ServerCookieCtxKey)
	if c == nil {
		return nil
	}

	return c.(*http.Cookie)
}

func ExtractCookie(r *http.Request) *http.Cookie {
	c, _ := r.Cookie(string(GetCookieKey()))
	return c
}

func CreateCookie(value string, options ...cookieOption) *http.Cookie {
	cookie := &http.Cookie{
		Name:     string(GetCookieKey()),
		Value:    value,
		Path:     "/",
		MaxAge:   10 * 60,
		HttpOnly: true,
		Secure:   true,
	}

	for _, opt := range options {
		opt(cookie)
	}

	return cookie
}

func DeleteCookie() *http.Cookie {
	c := &http.Cookie{
		Name:     string(GetCookieKey()),
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
		Secure:   true,
	}

	return c
}

func WithName(name string) cookieOption {
	return func(c *http.Cookie) {
		c.Name = name
	}
}

func WithMaxAge(seconds int) cookieOption {
	return func(c *http.Cookie) {
		c.MaxAge = seconds
	}
}

func WithValue(value string) cookieOption {
	return func(c *http.Cookie) {
		c.Value = value
	}
}
