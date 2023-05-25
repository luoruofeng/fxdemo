package srv

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Abc2 struct {
	logger  *zap.Logger
	Content string
}

func NewAbc2(lc fx.Lifecycle, logger *zap.Logger, content string) Abc2 {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("HAHA Abc2 Start building!!!")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("HAHA Abc2 begins destruction!!!")
			return nil
		},
	})

	return Abc2{logger: logger, Content: content}
}
