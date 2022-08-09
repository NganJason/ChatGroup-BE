package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/ChatGroup-BE/internal/utils"
	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
	"github.com/NganJason/ChatGroup-BE/pkg/clog"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func CreateSocketMiddleware(next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		clog.Info(r.Context(), "initiating socket")
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			clog.Error(
				r.Context(),
				fmt.Sprintf("socket handshake err=%s", err.Error()),
			)

			handleErr(
				next,
				w,
				r,
				cerr.New(
					fmt.Sprintf("init websocket err=%s", err.Error()),
					http.StatusBadGateway,
				),
			)

			return
		}

		ctx := context.WithValue(
			r.Context(),
			utils.SocketCtxKey,
			conn,
		)
		r = r.WithContext(ctx)

		next(w, r)
	}

	return fn
}
