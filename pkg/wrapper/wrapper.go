package wrapper

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/NganJason/ChatGroup-BE/internal/middleware"
	"github.com/NganJason/ChatGroup-BE/internal/utils"
	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
	"github.com/NganJason/ChatGroup-BE/pkg/clog"
	"github.com/NganJason/ChatGroup-BE/pkg/cookies"
	"github.com/gorilla/websocket"
)

type Processor func(ctx context.Context, req, resp interface{}) error

func WrapProcessor(
	proc Processor,
	req, resp interface{},
	needAuth bool,
	socket bool,
) http.HandlerFunc {
	if socket && needAuth {
		return middleware.CheckAuthMiddleware(
			middleware.CreateSocketMiddleware(
				Wrapper(proc, req, resp, socket),
			),
		)
	}
	if socket {
		return middleware.CreateSocketMiddleware(Wrapper(proc, req, resp, socket))
	}

	if needAuth {
		return middleware.CheckAuthMiddleware(Wrapper(proc, req, resp, socket))
	}

	return Wrapper(proc, req, resp, socket)
}

func Wrapper(
	proc Processor,
	req, resp interface{},
	socket bool,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newReq := reflect.New(reflect.TypeOf(req).Elem()).Interface()
		newResp := reflect.New(reflect.TypeOf(resp).Elem()).Interface()

		ctx := clog.ContextWithTraceID(
			r.Context(),
			strconv.FormatUint(uint64(time.Now().Unix()), 10),
		)

		err := cerr.GetErrFromCtx(ctx)
		if err != nil {
			writeToClient(ctx, w, newResp, err)
			return
		}

		err = json.NewDecoder(r.Body).Decode(&newReq)
		if err != nil {
			if err == io.EOF {
				newReq = nil
			} else {
				writeToClient(ctx, w, newResp, err)
				return
			}
		}

		ctx = cookies.InitServerCookie(ctx)

		err = proc(ctx, newReq, newResp)
		if err != nil {
			writeToClient(ctx, w, newResp, err)
			return
		}

		setCookie(ctx, w)

		if !socket {
			writeToClient(ctx, w, newResp, nil)
		}
	}
}

func writeToClient(ctx context.Context, w http.ResponseWriter, resp interface{}, err error) {
	setDebugMessage(ctx, resp, err)
	jsonResp, _ := json.Marshal(resp)

	socketConn := ctx.Value(utils.SocketCtxKey)
	if socketConn == nil {
		if err != nil {
			code := cerr.Code(err)
			w.WriteHeader(code)
		} else {
			w.WriteHeader(http.StatusOK)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
	} else {
		conn, _ := socketConn.(*websocket.Conn)
		conn.WriteMessage(websocket.BinaryMessage, jsonResp)
	}
}

func setDebugMessage(ctx context.Context, resp interface{}, err error) {
	if resp == nil || err == nil {
		return
	}

	msg := err.Error()

	debugMsgField := "DebugMsg"
	structField, found := reflect.TypeOf(resp).Elem().FieldByName(debugMsgField)
	if !found {
		clog.Error(ctx, "debug_msg not found")
		return
	}

	fieldType := structField.Type
	if fieldType.Kind() != reflect.Ptr || fieldType.Elem().Kind() != reflect.String {
		return
	}

	requiredField := reflect.ValueOf(resp).Elem().FieldByName(debugMsgField)

	if requiredField.CanSet() {
		var finalMsg string

		elem := requiredField.Elem()
		if elem.IsValid() && len(elem.String()) != 0 {
			finalMsg = elem.String() + ": " + msg
		} else {
			finalMsg = msg
		}

		requiredField.Set(reflect.ValueOf(&finalMsg))
	}
}

func setCookie(ctx context.Context, w http.ResponseWriter) {
	c := cookies.GetServerCookieFromCtx(ctx)
	if c == nil {
		return
	}

	http.SetCookie(w, c)
}
