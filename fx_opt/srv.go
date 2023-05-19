package fx_opt

import (
	"net/http"

	fxp "github.com/luoruofeng/fxdemo/fx_opt/component/provide"
	fxhttp "github.com/luoruofeng/fxdemo/fx_opt/component/provide/http"
	"github.com/luoruofeng/fxdemo/http/handler"
	"github.com/luoruofeng/fxdemo/http/middleware"

	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func GetApp() *fx.App {

	//handlers Provide
	handlerProv := fx.Provide(
		fxhttp.AllAsRoute(handler.NewEchoHandler, handler.NewHelloHandler)...,
	)

	//middlewares Provide
	middlewareProv := fx.Provide(
		fxhttp.AllAsMiddleware(middleware.NewLogMiddleware, middleware.NewLogMiddleware2)...,
	)

	//config Provide
	configAnno := fx.Annotate(fxp.NewConfig)
	cnfProv := fx.Provide(
		//config file path
		func() string {
			r := ""
			return r
		},
		configAnno,
	)

	//Http server Provide
	httpMuxAnno := fx.Annotate(
		fxhttp.NewHTTPRouter,
		fx.ParamTags(`group:"middlewares"`, `group:"handlers"`),
	)

	httpSrvProv := fx.Provide(
		httpMuxAnno,
		fxhttp.NewHTTPServer,
	)

	//logger Provide
	loggerProv := fx.Provide(fxp.NewLogger)

	app := fx.New(
		// The same Zap logger for Fx's own logs as well.
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),

		//Provides
		handlerProv,
		middlewareProv,
		loggerProv,
		cnfProv,
		httpSrvProv,

		//Invoke
		fx.Invoke(func(*mux.Router) {}, func(*http.Server) {}),
	)
	return app
}
