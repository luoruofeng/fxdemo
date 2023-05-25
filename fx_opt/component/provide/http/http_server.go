package http

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/luoruofeng/fxdemo/conf"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewHTTPServer(lc fx.Lifecycle, logger *zap.Logger, c *conf.Config, r *mux.Router) *http.Server {
	server := &http.Server{
		Addr:         c.HttpAddr,
		Handler:      r,
		WriteTimeout: time.Duration(c.HttpWriteOverTime) * time.Second,
		ReadTimeout:  time.Duration(c.HttpReadOverTime) * time.Second,
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting HTTP server!", zap.String("addr", server.Addr))
			ln, err := net.Listen("tcp", server.Addr)
			if err != nil {
				return err
			}
			go server.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping HTTP server!")
			return server.Shutdown(ctx)
		},
	})

	return server
}
