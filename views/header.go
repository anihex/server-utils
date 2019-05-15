package views

import (
	"net/http"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/anihex/server-utils/tools"
)

// SendHeader sends a header to the client. It also logs this call.
func SendHeader(w http.ResponseWriter, r *http.Request, Status int) {
	w.WriteHeader(Status)
	_, filename, line, _ := runtime.Caller(1)
	filename = filepath.Base(filename)

	RemoteAddr := tools.GetIP(r)

	if !strings.HasPrefix(r.RequestURI, "/status") && TEST_MODE == false {
		Logger.Printf("[%7s] [%s; Line %d] [Header] [%d] [%s] %s [Header Send]", r.Method, filename, line, Status, RemoteAddr, r.RequestURI)
	}
}
