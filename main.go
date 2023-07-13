package main

import (
	"github.com/luoruofeng/fxdemo/cmd"
	"github.com/luoruofeng/fxdemo/fx_opt"
)

func main() {
	configPathMap := cmd.GetConfigFilePath()
	fxSrv := fx_opt.NewFxSrv(configPathMap)
	fxSrv.Setup()
	fxSrv.Start()
	fxSrv.Shutddown()
}
