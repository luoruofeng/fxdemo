package fx_opt

import (
	"context"
	"log"
	"net/http"
	"time"

	fxp "github.com/luoruofeng/fxdemo/fx_opt/component/provide"
	fxhttp "github.com/luoruofeng/fxdemo/fx_opt/component/provide/http"
	"github.com/luoruofeng/fxdemo/http/handler"
	"github.com/luoruofeng/fxdemo/http/middleware"

	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewFxSrv(configPath string) FxSrv {
	return FxSrv{
		configFilePath: configPath,
	}
}

type FxSrv struct {
	app            *fx.App
	configFilePath string
}

func (f *FxSrv) Start() {
	err := f.app.Start(context.Background())
	if err != nil {
		panic(err)
	}
	<-f.app.Done()
}

func (f *FxSrv) Shutddown() {
	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := f.app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}

func (f *FxSrv) Setup() {

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
			return f.configFilePath
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
	f.app = app
}
