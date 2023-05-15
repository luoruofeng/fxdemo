package main

import (
	"net/http"

	component "github.com/luoruofeng/fxdemo/fx_component"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			component.NewHTTPServer,
			component.NewEchoHandler,
			component.NewServeMux,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
