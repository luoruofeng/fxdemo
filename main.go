package main

import (
	"github.com/luoruofeng/fxdemo/cmd"
	"github.com/luoruofeng/fxdemo/fx_opt"
)

func main() {
	fxSrv := fx_opt.NewFxSrv(cmd.GetConfigFilePath())
	fxSrv.Setup()
	fxSrv.Start()
	fxSrv.Shutddown()
}
