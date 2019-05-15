package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/anihex/server-utils/tools"
)

// Time adds the current time to the context
func Time(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var timer tools.TimeType = 1

		ctx := context.WithValue(
			r.Context(),
			timer,
			time.Now(),
		)

		f(w, r.WithContext(ctx))
	}
}
