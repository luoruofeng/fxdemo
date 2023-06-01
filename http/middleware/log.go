package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

// Define our struct
type LogMiddleware struct {
	logger *zap.Logger
}

func NewLogMiddleware(logger *zap.Logger) *LogMiddleware {
	return &LogMiddleware{logger: logger}
}

// Middleware function, which will be called for each request
func (l *LogMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.logger.Info("this is log middleware")
		next.ServeHTTP(w, r)

	})
}
