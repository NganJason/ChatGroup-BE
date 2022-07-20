package middleware

import (
	"net/http"

	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
)

func handleErr(next http.HandlerFunc, w http.ResponseWriter, r *http.Request, err error) {
	r = r.WithContext(cerr.AddErrToCtx(r.Context(), err))
	next(w, r)
}
