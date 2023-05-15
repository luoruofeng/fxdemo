package main

import (
	"net/http"

	f "github.com/luoruofeng/fxdemo/fx_component"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	// You can use the same Zap logger for Fx's own logs as well.
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			f.NewHTTPServer,
			f.AsRoute(f.NewEchoHandler),
			f.AsRoute(f.NewHelloHandler),
			f.NewLogger,
			fx.Annotate(
				f.NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
