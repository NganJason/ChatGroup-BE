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
	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
	"github.com/NganJason/ChatGroup-BE/pkg/clog"
	"github.com/NganJason/ChatGroup-BE/pkg/cookies"
)

type Processor func(ctx context.Context, req, resp interface{}) error

func WrapProcessor(
	proc Processor,
	req, resp interface{},
	needAuth bool,
) http.HandlerFunc {
	if needAuth {
		return middleware.CheckAuthMiddleware(Wrapper(proc, req, resp))
	}

	return Wrapper(proc, req, resp)
}

func Wrapper(
	proc Processor,
	req, resp interface{},
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

		writeToClient(ctx, w, newResp, nil)
	}
}

func writeToClient(ctx context.Context, w http.ResponseWriter, resp interface{}, err error) {
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		code := cerr.Code(err)
		w.WriteHeader(code)

		setDebugMessage(ctx, resp, err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
	}

	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func setDebugMessage(ctx context.Context, resp interface{}, msg string) {
	if resp == nil {
		return
	}

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
