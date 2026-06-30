package middlewares

import (
	"net/http"
	"time"

	"goblog/pkg/logger"

	"go.uber.org/zap"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (w *responseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

// CurlLog 记录入站 HTTP 请求（异步写入 curl channel）
func CurlLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(rw, r)

		logger.LogCurl(
			zap.String("type", "incoming"),
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("query", r.URL.RawQuery),
			zap.String("ip", r.RemoteAddr),
			zap.Int("status", rw.status),
			zap.Duration("elapsed", time.Since(start)),
		)
	})
}
