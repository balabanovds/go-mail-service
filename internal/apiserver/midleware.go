package apiserver

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ctxKey int8

const (
	ctxReqID ctxKey = iota
)

func jsonContent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		zap.L().Info("http_request",
			zapRequestID(r),
			zap.String("method", r.Method),
			zap.String("uri", r.RequestURI),
			zap.String("from", r.RemoteAddr),
		)
		start := time.Now()
		lrw := newLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)

		zap.L().Info("http_request",
			zapRequestID(r),
			zap.Int("code", lrw.statusCode),
			zap.Duration("duration", time.Since(start)),
		)
	})
}

func (s *ApiServer) withRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		ctx := context.WithValue(r.Context(), ctxReqID, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
