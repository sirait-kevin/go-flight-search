package middlewares

import (
	"context"
	"github.com/sirupsen/logrus"
	"go-flight-search/pkg/helper"
	"go-flight-search/pkg/logger"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger.InitLogger(true)
		logger := logger.Log.WithFields(logrus.Fields{
			"method":  r.Method,
			"path":    r.URL.Path,
			"latency": time.Since(start).String(),
		})

		ctx := context.WithValue(r.Context(), "logger", logger)
		rw := &helper.ResponseWriter{ResponseWriter: w, StatusCode: http.StatusOK}
		next.ServeHTTP(rw, r.WithContext(ctx))

		duration := time.Since(start)
		logger.WithFields(logrus.Fields{
			"status":   rw.StatusCode,
			"duration": duration,
		}).Info("request completed")
	})
}
