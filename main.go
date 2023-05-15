package main

import (
	"net/http"

	component "github.com/luoruofeng/fxdemo/fx_component"

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
			component.NewHTTPServer,
			fx.Annotate(
				component.NewEchoHandler,
				fx.As(new(component.Route)),
			),
			component.NewServeMux,
			// zap.NewExample,
			//使用自定义对的logger
			component.NewLogger,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
