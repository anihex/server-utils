package views

import (
	"net/http"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/anihex/json"
	"github.com/anihex/server-utils/tools"
)

// SendJSON takes an Object and encodes it to JSON. The result will then be
// send to the Client.
func SendJSON(w http.ResponseWriter, r *http.Request, o interface{}, Status int) {
	data, err := json.Marshal(o)
	if err != nil {
		data, _ = json.Marshal(tools.H{"error": "Internal Sever Error"})
		return
	}

	send(w, r, data, "application/json; charset=utf-8", Status)
}

// send takes a byte Array and sends it to the client using the given
// Content-Type and the given Status-Code.
func send(w http.ResponseWriter, r *http.Request, data []byte, ContentType string, Status int) {
	RemoteAddr := tools.GetIP(r)
	_, filename, line, _ := runtime.Caller(2)
	filename = filepath.Base(filename)

	w.Header().Set("Content-Type", ContentType)
	w.WriteHeader(Status)

	timeElapsed := getTime(r)

	w.Write(data)

	if !strings.HasPrefix(r.RequestURI, "/status") && TEST_MODE == false {
		Logger.Printf(
			"[%7s] [%s; Line %d] [JSON] [%d] [%s] [%v] %s\n",
			r.Method,
			filename,
			line,
			Status,
			RemoteAddr,
			timeElapsed,
			r.RequestURI,
		)
	}
}
