package views

import (
	"net/http"
	"strings"
)

// SendBytes sends bytes over HTTP. It uses the given content-type and the
// given HTTP Status-Code
func SendBytes(w http.ResponseWriter, r *http.Request, data []byte, ContentType string, Status int) {
	w.Header().Set("Content-Type", ContentType)
	w.WriteHeader(Status)
	w.Write(data)

	if !strings.HasPrefix(r.RequestURI, "/status") && TEST_MODE == false {
		Logger.Printf("[%7s] [%d] %s\n [Bytes Send]", r.Method, Status, r.RequestURI)
	}
}
