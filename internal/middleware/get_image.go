package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/NganJason/ChatGroup-BE/internal/utils"
	"github.com/NganJason/ChatGroup-BE/pkg/cerr"
)

func GetImageMiddleware(next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		paths, ok := r.URL.Query()["path"]
		if !ok {
			handleErr(
				next,
				w,
				r,
				cerr.New(
					"query path error",
					http.StatusBadGateway,
				),
			)
			return
		}

		imgPath := getFilePath(paths[0])

		fileByte, err := ioutil.ReadFile(imgPath)
		if err != nil {
			handleErr(
				next,
				w,
				r,
				cerr.New(
					fmt.Sprintf(
						"read file path err=%s",
						err.Error(),
					),
					http.StatusBadGateway,
				),
			)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(fileByte)
	}

	return fn
}

func getFilePath(path string) string {
	return fmt.Sprintf("%s/%s", utils.ImageDir, path)
}
