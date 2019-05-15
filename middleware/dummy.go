package middleware

import (
	"net/http"
)

// Dummy is a simple dummy middleware
func Dummy(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
	}
}
