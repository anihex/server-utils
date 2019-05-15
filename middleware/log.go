package middleware

import (
	"net/http"
	"strings"

	"github.com/anihex/server-utils/tools"
)

// Log logs the requests. It uses the package logger to log.
func Log(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.RequestURI, "/status") {
			ip := tools.GetIP(r)

			lg.Printf("[%7s] [%s] %s", r.Method, ip, r.RequestURI)
		}

		f(w, r)
	}
}
