package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

// Define our struct
type LogMiddleware2 struct {
	logger *zap.Logger
}

func NewLogMiddleware2(logger *zap.Logger) *LogMiddleware2 {
	return &LogMiddleware2{logger: logger}
}

// Middleware function, which will be called for each request
func (l *LogMiddleware2) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.logger.Info("This is log2 middleware")
		next.ServeHTTP(w, r)

	})
}
