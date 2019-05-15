package views

import (
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/anihex/server-utils/tools"
)

func fileExists(Filename string) bool {
	stat, err := os.Stat(Filename)
	if err != nil {
		return false
	}

	if stat.IsDir() {
		return false
	}

	return true
}

// SendFile sends the given file to the client if the file was found and can be
// accessed. Otherwise it will send a "Not Found".
func SendFile(w http.ResponseWriter, r *http.Request, FileName string) {
	_, filename, line, _ := runtime.Caller(1)
	filename = filepath.Base(filename)

	RemoteAddr := tools.GetIP(r)

	if !fileExists(FileName) {
		w.WriteHeader(http.StatusNotFound)
		Logger.Printf("[%7s] [%s; Line %d] [File] [%d] [%s] %s", r.Method, filename, line, http.StatusNotFound, RemoteAddr, r.RequestURI)
	} else {
		Logger.Printf("[%7s] [%s; Line %d] [File] [%d] [%s] %s", r.Method, filename, line, http.StatusOK, RemoteAddr, r.RequestURI)
		http.ServeFile(w, r, FileName)
	}
}
