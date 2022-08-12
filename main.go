package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/NganJason/ChatGroup-BE/internal/config"
	"github.com/NganJason/ChatGroup-BE/internal/processor"
	"github.com/NganJason/ChatGroup-BE/pkg/clog"
	"github.com/NganJason/ChatGroup-BE/pkg/socket"
	"github.com/NganJason/ChatGroup-BE/pkg/wrapper"
	"github.com/rs/cors"
)

func main() {
	ctx := context.Background()
	mux := http.NewServeMux()
	config.InitConfig()
	socket.InitHub()

	for _, proc := range processor.GetAllProcessors() {
		mux.HandleFunc(
			proc.Path,
			wrapper.WrapProcessor(
				proc.Processor,
				proc.Req,
				proc.Resp,
				proc.NeedAuth,
				proc.Socket,
				proc.File,
				proc.GetFile,
			),
		)
	}

	clog.SetMinLogLevel(clog.LevelInfo)
	c := cors.New(
		cors.Options{
			AllowedOrigins:   []string{"https://delicate-queijadas-d1ee5b.netlify.app", "http://localhost:3000"},
			AllowCredentials: true,
			AllowedMethods:   []string{"POST", "GET", "OPTIONS"},
		},
	)
	handler := c.Handler(mux)

	clog.Info(ctx, fmt.Sprintf("Listening to port %s", GetPort()))
	err := http.ListenAndServe(GetPort(), handler)
	if err != nil {
		clog.Fatal(ctx, fmt.Sprintf("error init server, %s", err.Error()))
	}
}

func GetPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "8082"
	}

	return ":" + port
}
