package middleware

import (
	"log/slog"
	"net/http"

	"user-service/internal/metrics"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func (r *statusRecorder) Write(b []byte) (int, error) {
	if r.status == 0 {
		r.status = http.StatusOK
	}
	return r.ResponseWriter.Write(b)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ip       = r.RemoteAddr
			protocol = r.Proto
			method   = r.Method
			path     = r.URL.Path
		)

		// increment request counter for metrics
		metrics.IncRequests()

		recorder := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(recorder, r)
		slog.Info("Incoming request",
			slog.String("ip", ip),
			slog.String("protocol", protocol),
			slog.String("method", method),
			slog.String("path", path),
			slog.Int("status_code", recorder.status),
		)
	})
}
