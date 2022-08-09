package processor

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NganJason/ChatGroup-BE/internal/utils"
	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
	"github.com/NganJason/ChatGroup-BE/pkg/cookies"
	"github.com/NganJason/ChatGroup-BE/pkg/socket"
	"github.com/NganJason/ChatGroup-BE/vo"
	"github.com/gorilla/websocket"
)

func CreateSocketProcessor(
	ctx context.Context,
	req,
	resp interface{},
) error {
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

	socketConn := ctx.Value(utils.SocketCtxKey)
	if socketConn == nil {
		return cerr.New(
			"socket connection not found",
			http.StatusBadGateway,
		)
	}

	conn, ok := socketConn.(*websocket.Conn)
	if !ok {
		return cerr.New(
			"connection assertion failed",
			http.StatusBadGateway,
		)
	}

	socket.GetHub().RegisterClient(
		userID,
		conn,
	)

	socket.GetHub().Broadcast(
		userID,
		[]uint64{userID},
		vo.SocketMessage{
			EventType: uint32(vo.ServerEvent),
			Message:   "handshake succeeded!",
		},
	)

	return nil
}
