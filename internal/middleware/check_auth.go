package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/NganJason/ChatGroup-BE/internal/handler"
	"github.com/NganJason/ChatGroup-BE/internal/model"
	"github.com/NganJason/ChatGroup-BE/internal/utils"
	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
	"github.com/NganJason/ChatGroup-BE/pkg/cookies"
)

func CheckAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c := cookies.ExtractCookie(r)
		if c == nil {
			err := cerr.New(
				"cookies not found",
				http.StatusUnauthorized,
			)
			handleErr(next, w, r, err)
			return
		}

		jwt := c.Value
		auth, err := utils.ParseJWTToken(jwt)
		if err != nil || auth == nil {
			err = cerr.New(
				fmt.Sprintf("parse jwt token err=%s", err.Error()),
				http.StatusBadGateway,
			)
			handleErr(next, w, r, err)
			return
		}

		if auth.Valid() != nil {
			err = cerr.New(
				fmt.Sprintf("token is not valid err=%s", auth.Valid().Error()),
				http.StatusUnauthorized,
			)

			handleErr(next, w, r, err)
			return
		}

		userIDStr := auth.Value
		userID, err := strconv.ParseUint(userIDStr, 0, 64)
		if err != nil {
			handleErr(
				next,
				w,
				r,
				cerr.New(
					fmt.Sprintf("parse userIDStr err=%s", err.Error()),
					http.StatusBadGateway,
				),
			)
			return
		}

		userDM, err := model.NewUserDM(r.Context())
		if err != nil {
			handleErr(next, w, r, err)
			return
		}

		h := handler.NewAuthHandler(r.Context(), userDM)

		isAuth, err := h.ValidateUser(&userID)
		if err != nil {
			handleErr(next, w, r, err)
			return
		}

		if !isAuth {
			err = cerr.New(
				"user is not authenticated",
				http.StatusUnauthorized,
			)
			handleErr(next, w, r, err)
			return
		}

		r = r.WithContext(cookies.AddClientCookieValToCtx(r.Context(), &userIDStr))

		next(w, r)
	}

	return fn
}
