package middleware

import (
	"net/http"
	"strconv"
)

// CORS adds CORS headers to a request
func CORS(Origins string, MaxAge int, Credentials bool, Methods string) func(f http.HandlerFunc) http.HandlerFunc {
	cred := "false"
	if Credentials {
		cred = "true"
	}
	mage := strconv.Itoa(MaxAge)

	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", Origins)
			w.Header().Set("Access-Control-Max-Age", mage)
			w.Header().Set("Access-Control-Allow-Credentials", cred)
			w.Header().Set("Access-Control-Allow-Methods", Methods)
			w.Header().Set("Access-Control-Allow-Headers", "*")

			f(w, r)
		}
	}
}
