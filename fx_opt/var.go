package fx_opt

import (
	"github.com/luoruofeng/fxdemo/srv"
	"go.uber.org/fx"
)

// 添加其他需要在fx中构建的实例的方法
var ConstructorFuncs = []interface{}{
	srv.NewAbc,

	fx.Annotate(
		func() string {
			return "这是Abc2的Content参数"
		},
		fx.ResultTags(`name:"abc2content"`),
	),

	fx.Annotate(
		srv.NewAbc2,
		fx.ParamTags(``, ``, `name:"abc2content"`),
	),

	srv.NewAbc3,
}

// 在ConstructorFuncs添加了方法后，如果需要在方法的参数中传递fx.Lifecycle，已实现fx.Hook。需要在下方添加fx的invoke方法。
var InvokeFuncs = []interface{}{
	func(srv.Abc) {},
	func(srv.Abc2) {},
	func(srv.Abc3) {},
}
