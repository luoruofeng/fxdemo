package srv

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Abc3 struct {
	logger *zap.Logger
	Abc2   Abc2
}

func NewAbc3(lc fx.Lifecycle, logger *zap.Logger, abc2 Abc2) Abc3 {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Abc3开始构建", zap.Any("abc2", abc2))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Abc3开始销毁")
			return nil
		},
	})

	return Abc3{logger: logger, Abc2: abc2}
}
