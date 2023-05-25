package srv

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Abc struct {
	logger *zap.Logger
}

func NewAbc(lc fx.Lifecycle, logger *zap.Logger) Abc {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("HAHA Abc Start building!!!")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("HAHA Abc begins destruction!!!")
			return nil
		},
	})

	return Abc{logger: logger}
}
