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
// It takes a request and increases the "toSkip" information by 1.
func prepContext(r *http.Request) *http.Request {
	skip := getSkipFile(r)
	res := r.WithContext(
		context.WithValue(r.Context(), toSkip, skip+1),
	)

	return res
}

func sendError(w http.ResponseWriter, r *http.Request, err error, status int, data []byte) {
	skip := getSkipFile(r)
	elapsedTime := getTime(r)

	_, filename, line, _ := runtime.Caller(skip + 1)
	filename = filepath.Base(filename)

	RemoteAddr := tools.GetIP(r)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(data)

	if !strings.HasPrefix(r.RequestURI, "/status") && TEST_MODE == false {
		Logger.Printf(
			"[%7s] [%s; Line %d] [JSON] [%d] [%s] [%v] %s (%s)",
			r.Method,
			filename,
			line,
			status,
			RemoteAddr,
			elapsedTime,
			r.RequestURI,
			err,
		)
	}
}

// AccessDeniedWithErr sends an error message with "Forbidden" as it's
// status code. It also sends a JSON Object with the error-message
// "ERR_FORBIDDEN".
// The error message will be displayed in the log.
func AccessDeniedWithErr(w http.ResponseWriter, r *http.Request, err error) {
	data := []byte(`{ "error": "ERR_FORBIDDEN" }`)

	sendError(w, r, err, http.StatusForbidden, data)
}

// ErrAccessDenied sends an error message with "Forbidden" as it's status code.
// It also sends a JSON Object with the error-message "ERR_FORBIDDEN".
// It uses the default error message for "Forbidden".
func ErrAccessDenied(w http.ResponseWriter, r *http.Request) {
	AccessDeniedWithErr(w, r, errors.New("Forbidden"))
}

// AccessDeniedIfErr send an ERR_FORBIDDEN to the client IF the passed err is
// not nil. In this case error will be placed into the context and logged.
// If a Response was send, the result will be true to indicate, that no further
// request handling is necessary.
func AccessDeniedIfErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		res := prepContext(r)
		AccessDeniedWithErr(w, res, err)
		return true
	}

	return false
}

// UnauthorizedWithErr sends an error message with "Unauthorized" as it's
// status code. It also sends a JSON Object with the error-message
// "ERR_UNAUTHORIZED".
// The error message will be displayed in the log.
func UnauthorizedWithErr(w http.ResponseWriter, r *http.Request, err error) {
	data := []byte(`{ "error": "ERR_UNAUTHORIZED" }`)

	sendError(w, r, err, http.StatusUnauthorized, data)
}

// ErrUnauthorized sends an error message with "Unauthorized" as it's status code.
// It also sends a JSON Object with the error-message "ERR_UNAUTHORIZED".
// It uses the default error message for "Unauthorized".
func ErrUnauthorized(w http.ResponseWriter, r *http.Request) {
	UnauthorizedWithErr(w, r, errors.New("Unauthorized"))
}

// UnauthorizedIfErr send an ERR_UNAUTHORIZED to the client IF the passed err is
// not nil. In this case error will be placed into the context and logged.
// If a Response was send, the result will be true to indicate, that no further
// request handling is necessary.
func UnauthorizedIfErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		res := prepContext(r)
		UnauthorizedWithErr(w, res, err)
		return true
	}

	return false
}

// NotFoundWithErr sends an error message with "Not Found" as it's status code.
// It also sends a JSON Object with the error-message "ERR_NOT_FOUND".
// The error message will be displayed in the log.
func NotFoundWithErr(w http.ResponseWriter, r *http.Request, err error) {
	data := []byte(`{ "error": "ERR_NOT_FOUND" }`)

	sendError(w, r, err, http.StatusNotFound, data)
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
		res := prepContext(r)
		NotFoundWithErr(w, res, err)
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

	sendError(w, r, err, http.StatusInternalServerError, data)
}

// ErrServerError sends an error message with "Internal Server Error" as it's
// status code.
// It also sends a JSON Object with the error-message "ERR_SERVER_ERROR".
// It uses the default error message for "Internal Server Error".
func ErrServerError(w http.ResponseWriter, r *http.Request) {
	ServerErrorWithErr(w, r, errors.New("Internal Server Error"))
}

// ServerErrorIfErr send an ERR_SERVER_ERROR to the client IF the passed err is
// not nil. In this case error will be placed into the context and logged.
// If a Response was send, the result will be true to indicate, that no further
// request handling is necessary.
func ServerErrorIfErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		res := prepContext(r)
		ServerErrorWithErr(w, res, err)
		return true
	}

	return false
}

// InvalidDataWithErr sends an error message with "Unprocessable Entity" as it's
// status code.
// It also sends a JSON Object with the error-message "ERR_INVALID_DATA".
// The error message will be displayed in the log.
func InvalidDataWithErr(w http.ResponseWriter, r *http.Request, err error) {
	data := []byte(`{ "error": "ERR_INVALID_DATA" }`)

	sendError(w, r, err, http.StatusUnprocessableEntity, data)
}

// ErrInvalidData sends an error message with "Unprocessable Entity" as it's
// status code.
// It also sends a JSON Object with the error-message "ERR_INVALID_DATA".
// It uses the default error message for "Unprocessable Entity".
func ErrInvalidData(w http.ResponseWriter, r *http.Request) {
	InvalidDataWithErr(w, r, errors.New("Unprocessable Entity"))
}

// InvalidDataIfErr send an ERR_INVALID_DATA to the client IF the passed err is
// not nil. In this case error will be placed into the context and logged.
// If a Response was send, the result will be true to indicate, that no further
// request handling is necessary.
func InvalidDataIfErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		res := prepContext(r)
		InvalidDataWithErr(w, res, err)
		return true
	}

	return false
}

// InvalidMediaTypeWithErr sends an error message with "Unsupported Media Type" as
// it's status code.
// It also sends a JSON Object with the error-message "ERR_UNSUPPORTED_MEDIA_TYPE".
// The error message will be displayed in the log.
func InvalidMediaTypeWithErr(w http.ResponseWriter, r *http.Request, err error) {
	data := []byte(`{ "error": "ERR_UNSUPPORTED_MEDIA_TYPE" }`)

	sendError(w, r, err, http.StatusUnsupportedMediaType, data)
}

// ErrInvalidMediaType sends an error message with "Unsupported Media Type" as
// it's status code.
// It also sends a JSON Object with the error-message "ERR_UNSUPPORTED_MEDIA_TYPE".
// It uses the default error message for "Unsupported Media Type".
func ErrInvalidMediaType(w http.ResponseWriter, r *http.Request) {
	InvalidMediaTypeWithErr(w, r, errors.New("Unsupported Media Type"))
}

// InvalidMediaTypeIfErr send an ERR_UNSUPPORTED_MEDIA_TYPE to the client IF
// the passed err is not nil. In this case error will be placed into the
// context and logged. If a Response was send, the result will be true to
// indicate, that no further request handling is necessary.
func InvalidMediaTypeIfErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		res := prepContext(r)
		InvalidMediaTypeWithErr(w, res, err)
		return true
	}

	return false
}

// BadRequestWithErr sends an error message with "Bad Request" as it's status code.
// It also sends a JSON Object with the error-message "ERR_BAD_REQUEST".
// The error message will be displayed in the log.
func BadRequestWithErr(w http.ResponseWriter, r *http.Request, err error) {
	data := []byte(`{ "error": "ERR_BAD_REQUEST" }`)

	sendError(w, r, err, http.StatusBadRequest, data)
}

// ErrBadRequest sends an error message with "Bad Request" as it's status code.
// It also sends a JSON Object with the error-message "ERR_BAD_REQUEST".
// It uses the default error message for "Bad Request".
func ErrBadRequest(w http.ResponseWriter, r *http.Request) {
	BadRequestWithErr(w, r, errors.New("Bad Request"))
}

// BadRequestIfErr send an ERR_BAD_REQUEST to the client IF the passed err is
// not nil. In this case error will be placed into the context and logged.
// If a Response was send, the result will be true to indicate, that no further
// request handling is necessary.
func BadRequestIfErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		res := prepContext(r)
		BadRequestWithErr(w, res, err)
		return true
	}

	return false
}
