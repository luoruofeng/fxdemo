package main

import (
	"context"

	"github.com/luoruofeng/fxdemo/fx_opt"
)

func main() {
	app := fx_opt.GetApp()
	defer app.Stop(context.Background())
	app.Run()
}
