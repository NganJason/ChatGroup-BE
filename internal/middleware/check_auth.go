package middleware

import (
	"net/http"
)

func CheckAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {

		next(w, r)
	}

	return fn
}
