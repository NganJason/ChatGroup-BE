package middleware

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/NganJason/ChatGroup-BE/internal/utils"
	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
)

func ParseFileMiddleware(next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 << 20)
		file, _, err := r.FormFile(utils.ImageFormKey)
		if err != nil {
			handleErr(
				next,
				w,
				r,
				cerr.New(
					fmt.Sprintf("retrieve file err=%s", err.Error()),
					http.StatusBadGateway,
				),
			)

			return
		}

		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			handleErr(
				next,
				w,
				r,
				cerr.New(
					fmt.Sprintf("read file err=%s", err.Error()),
					http.StatusBadGateway,
				),
			)

			return
		}

		ctx := context.WithValue(
			r.Context(),
			utils.ImageCtxKey,
			fileBytes,
		)
		r = r.WithContext(ctx)

		next(w, r)
	}

	return fn
}
