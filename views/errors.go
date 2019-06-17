package views

import (
	"context"
	"errors"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/anihex/server-utils/tools"
)

// getSkipFile checks how many files need to be skipped for the log.
// "Files" refer to source code files.
func getSkipFile(r *http.Request) int {
	result := 1
	if ctx := r.Context().Value(toSkip); ctx != nil {
		skip, ok := ctx.(int)
		if ok && skip > 0 {
			result = skip
		}
	}

	return result
}

// getTime reads the starttime from the Request-Context and calcs how much time
// elapsed since then.
func getTime(r *http.Request) time.Duration {
	if ctx := r.Context().Value(toTime); ctx != nil {
		value, ok := ctx.(time.Time)
		if ok {
			return time.Since(value)
		}
	}

	return 0
}

// getError reads the error message from the context if available. If not, the
// default value will be used.
func getErr(r *http.Request, def string) string {
	if ctx := r.Context().Value(gotError); ctx != nil {
		value, ok := ctx.(string)
		if ok {
			return value
		}
	}

	return def
}

// preContext prepares the context for later use.
// It takes a request and an error and checks if the error was nil.
// If the error is not nil, the error gets stored inside the context and
// the "skipFile" gets increased by one so that the final views can handle this.
func prepContext(r *http.Request, err error) (*http.Request, bool) {
	if err == nil {
		return nil, false
	}

	skip := getSkipFile(r)
	res := r.WithContext(
		context.WithValue(r.Context(), toSkip, skip+1),
	)

	res = res.WithContext(
		context.WithValue(res.Context(), gotError, err.Error()),
	)

	return res, true
}

// AccessDeniedWithErr sends an error message with "Forbidden" as it's
// status code. It also sends a JSON Object with the error-message
// "ERR_FORBIDDEN".
// The error message will be displayed in the log.
func AccessDeniedWithErr(w http.ResponseWriter, r *http.Request, err error) {
	data := []byte(`{ "error": "ERR_FORBIDDEN" }`)

	skip := getSkipFile(r)
	elapsedTime := getTime(r)

	_, filename, line, _ := runtime.Caller(skip)
	filename = filepath.Base(filename)

	RemoteAddr := tools.GetIP(r)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusForbidden)
	w.Write(data)

	if !strings.HasPrefix(r.RequestURI, "/status") && TEST_MODE == false {
		Logger.Printf(
			"[%7s] [%s; Line %d] [JSON] [%d] [%s] [%v] %s (%s)",
			r.Method,
			filename,
			line,
			http.StatusForbidden,
			RemoteAddr,
			elapsedTime,
			r.RequestURI,
			err,
		)
	}
}

// ErrAccessDenied sends an error message with "Forbidden" as it's status code.
// It also sends a JSON Object with the error-message "ERR_FORBIDDEN".
// If the context of the request contains an error message, it will be displayed
// instead of the regular "Forbidden".
func ErrAccessDenied(w http.ResponseWriter, r *http.Request) {
	AccessDeniedWithErr(w, r, errors.New("Forbidden"))
}

// AccessDeniedIfErr send an ERR_FORBIDDEN to the client IF the passed err is
// not nil. In this case error will be placed into the context and logged.
// If a Response was send, the result will be true to indicate, that no further
// request handling is necessary.
func AccessDeniedIfErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		prepContext(r, err)
		AccessDeniedWithErr(w, r, err)
		return true
	}

	return false
}

// NotFoundWithErr sends an error message with "Not Found" as it's status code.
// It also sends a JSON Object with the error-message "ERR_NOT_FOUND".
// The error message will be displayed in the log.
func NotFoundWithErr(w http.ResponseWriter, r *http.Request, err error) {
	data := []byte(`{ "error": "ERR_NOT_FOUND" }`)

	skip := getSkipFile(r)
	elapsedTime := getTime(r)

	_, filename, line, _ := runtime.Caller(skip)
	filename = filepath.Base(filename)

	RemoteAddr := tools.GetIP(r)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	w.Write(data)

	if !strings.HasPrefix(r.RequestURI, "/status") && TEST_MODE == false {
		Logger.Printf(
			"[%7s] [%s; Line %d] [JSON] [%d] [%s] [%v] %s (%s)",
			r.Method,
			filename,
			line,
			http.StatusNotFound,
			RemoteAddr,
			elapsedTime,
			r.RequestURI,
			err,
		)
	}
}

// ErrNotFound sends an error message with "Not Found" as it's status code.
// It also sends a JSON Object with the error-message "ERR_NOT_FOUND".
// It uses the default error message for "Not Found".
func ErrNotFound(w http.ResponseWriter, r *http.Request) {
	NotFoundWithErr(w, r, errors.New("Not Found"))
}

// NotFoundIfErr send an ERR_NOT_FOUND to the client IF the passed err is
// not nil. In this case error will be placed into the context and logged.
// If a Response was send, the result will be true to indicate, that no further
// request handling is necessary.
func NotFoundIfErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		prepContext(r, err)
		NotFoundWithErr(w, r, err)
		return true
	}

	return false
}

// ServerErrorWithErr sends an error message with "Internal Server Error" as it's
// status code.
// It also sends a JSON Object with the error-message "ERR_SERVER_ERROR".
// The error message will be displayed in the log.
func ServerErrorWithErr(w http.ResponseWriter, r *http.Request, err error) {
	data := []byte(`{ "error": "ERR_SERVER_ERROR" }`)

	skip := getSkipFile(r)
	elapsedTime := getTime(r)

	_, filename, line, _ := runtime.Caller(skip)
	filename = filepath.Base(filename)

	RemoteAddr := tools.GetIP(r)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(data)

	if !strings.HasPrefix(r.RequestURI, "/status") && TEST_MODE == false {
		Logger.Printf(
			"[%7s] [%s; Line %d] [JSON] [%d] [%s] [%v] %s (%s)",
			r.Method,
			filename,
			line,
			http.StatusInternalServerError,
			RemoteAddr,
			elapsedTime,
			r.RequestURI,
			err,
		)
	}
}

// ErrServerError sends an error message with "Internal Server Error" as it's
// status code.
// It also sends a JSON Object with the error-message "ERR_SERVER_ERROR".
// If the context of the request contains an error message, it will be displayed
// instead of the regular "Forbidden".
func ErrServerError(w http.ResponseWriter, r *http.Request) {
	ServerErrorWithErr(w, r, errors.New("Internal Server Error"))
}

// ServerErrorIfErr send an ERR_SERVER_ERROR to the client IF the passed err is
// not nil. In this case error will be placed into the context and logged.
// If a Response was send, the result will be true to indicate, that no further
// request handling is necessary.
func ServerErrorIfErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		prepContext(r, err)
		ServerErrorWithErr(w, r, err)
		return true
	}

	return false
}

// ErrInvalidData sends an error message with "Unprocessable Entity" as it's
// status code.
// It also sends a JSON Object with the error-message "ERR_INVALID_DATA".
// If the context of the request contains an error message, it will be displayed
// instead of the regular "Forbidden".
func ErrInvalidData(w http.ResponseWriter, r *http.Request) {
	data := []byte(`{ "error": "ERR_INVALID_DATA" }`)

	skip := getSkipFile(r)
	err := getErr(r, "Unprocessable Entity")
	elapsedTime := getTime(r)

	_, filename, line, _ := runtime.Caller(skip)
	filename = filepath.Base(filename)

	RemoteAddr := tools.GetIP(r)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(data)

	if !strings.HasPrefix(r.RequestURI, "/status") && TEST_MODE == false {
		Logger.Printf(
			"[%7s] [%s; Line %d] [JSON] [%d] [%s] [%v] %s (%s)",
			r.Method,
			filename,
			line,
			http.StatusUnprocessableEntity,
			RemoteAddr,
			elapsedTime,
			r.RequestURI,
			err,
		)
	}
}

// InvalidDataIfErr send an ERR_INVALID_DATA to the client IF the passed err is
// not nil. In this case error will be placed into the context and logged.
// If a Response was send, the result will be true to indicate, that no further
// request handling is necessary.
func InvalidDataIfErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if res, ok := prepContext(r, err); ok {
		ErrInvalidData(w, res)
		return true
	}

	return false
}

// ErrInvalidMediaType sends an error message with "Unsupported Media Type" as
// it's status code.
// It also sends a JSON Object with the error-message "ERR_UNSUPPORTED_MEDIA_TYPE".
// If the context of the request contains an error message, it will be displayed
// instead of the regular "Forbidden".
func ErrInvalidMediaType(w http.ResponseWriter, r *http.Request) {
	data := []byte(`{ "error": "ERR_UNSUPPORTED_MEDIA_TYPE" }`)

	skip := getSkipFile(r)
	err := getErr(r, "Unsupported Media Type")
	elapsedTime := getTime(r)

	_, filename, line, _ := runtime.Caller(skip)
	filename = filepath.Base(filename)

	RemoteAddr := tools.GetIP(r)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusUnsupportedMediaType)
	w.Write(data)

	if !strings.HasPrefix(r.RequestURI, "/status") && TEST_MODE == false {
		Logger.Printf(
			"[%7s] [%s; Line %d] [JSON] [%d] [%s] [%v] %s (%s)",
			r.Method,
			filename,
			line,
			http.StatusUnsupportedMediaType,
			RemoteAddr,
			elapsedTime,
			r.RequestURI,
			err,
		)
	}
}

// InvalidMediaTypeIfErr send an ERR_UNSUPPORTED_MEDIA_TYPE to the client IF
// the passed err is not nil. In this case error will be placed into the
// context and logged. If a Response was send, the result will be true to
// indicate, that no further request handling is necessary.
func InvalidMediaTypeIfErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if res, ok := prepContext(r, err); ok {
		ErrInvalidMediaType(w, res)
		return true
	}

	return false
}

// ErrBadRequest sends an error message with "Bad Request" as it's status code.
// It also sends a JSON Object with the error-message "ERR_BAD_REQUEST".
// If the context of the request contains an error message, it will be displayed
// instead of the regular "Forbidden".
func ErrBadRequest(w http.ResponseWriter, r *http.Request) {
	data := []byte(`{ "error": "ERR_BAD_REQUEST" }`)

	skip := getSkipFile(r)
	err := getErr(r, "Bad Request")
	elapsedTime := getTime(r)

	_, filename, line, _ := runtime.Caller(skip)
	filename = filepath.Base(filename)

	RemoteAddr := tools.GetIP(r)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(data)

	if !strings.HasPrefix(r.RequestURI, "/status") && TEST_MODE == false {
		Logger.Printf(
			"[%7s] [%s; Line %d] [JSON] [%d] [%s] [%v] %s (%s)",
			r.Method,
			filename,
			line,
			http.StatusBadRequest,
			RemoteAddr,
			elapsedTime,
			r.RequestURI,
			err,
		)
	}
}

// BadRequestIfErr send an ERR_BAD_REQUEST to the client IF the passed err is
// not nil. In this case error will be placed into the context and logged.
// If a Response was send, the result will be true to indicate, that no further
// request handling is necessary.
func BadRequestIfErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if res, ok := prepContext(r, err); ok {
		ErrBadRequest(w, res)
		return true
	}

	return false
}
