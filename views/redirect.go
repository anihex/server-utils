package views

import (
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/anihex/server-utils/tools"
)

// Redirect send a Redirect-Header to the Client. If "Perm" is false it sends a
// temporary redirect. The "URL" is the target of the redirect.
func Redirect(w http.ResponseWriter, r *http.Request, Perm bool, URL string) {
	StatusCode := 303
	if Perm {
		StatusCode = 301
	}

	_, filename, line, _ := runtime.Caller(1)
	filename = filepath.Base(filename)

	RemoteAddr := tools.GetIP(r)

	w.Header().Set("Status", strconv.Itoa(StatusCode))
	w.Header().Set("Location", URL)
	w.WriteHeader(StatusCode)

	if !strings.HasPrefix(r.RequestURI, "/status") && TEST_MODE == false {
		Logger.Printf("[%7s] [%s; Line %d] [%d] [%s] %s -> %s [Header Send]", r.Method, filename, line, StatusCode, RemoteAddr, r.RequestURI, URL)
	}
}
