package apiserver

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (s *ApiServer) clientError(w http.ResponseWriter, r *http.Request, code int, err error) {
	zap.L().Warn("client error",
		zap.Int("code", code),
		zapRequestID(r),
		zap.String("method", r.Method),
		zap.String("uri", r.RequestURI),
		zap.String("from", r.RemoteAddr),
		zap.Error(err),
	)
	s.respond(w, code, map[string]string{"error": err.Error()})
}

func (s *ApiServer) serverError(w http.ResponseWriter, err error) {
	zap.L().Error("server error",
		zap.Error(err),
	)
	s.respond(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
}

func (s *ApiServer) respond(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}

func zapRequestID(r *http.Request) zap.Field {
	requestID, ok := r.Context().Value(ctxReqID).(string)
	if !ok {
		return zap.String("request_id", "")
	}
	return zap.String("request_id", requestID)
}
