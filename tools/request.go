package tools

import (
	"net/http"
	"strings"
)

// GetIP reads the IP from the reuqest and it's headers. It returns the "real"
// IP as best as it can.
func GetIP(r *http.Request) (result string) {
	result = r.RemoteAddr

	remoteAddr := r.RemoteAddr
	realIP := strings.TrimSpace(r.Header.Get("X-Real-IP"))
	forwardedIP := strings.TrimSpace(r.Header.Get("X-Forwarded-For"))

	if remoteAddr != "" {
		result = remoteAddr
	}

	if realIP != "" {
		result = realIP
	}

	if forwardedIP != "" {
		result = forwardedIP
	}

	return
}
