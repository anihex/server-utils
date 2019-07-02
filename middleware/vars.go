package middleware

import (
	"log"

	"github.com/anihex/server-utils/tools"
)

var lg *log.Logger = tools.DefaultLogger

// SetLog sets the package logger
func SetLog(logger *log.Logger) {
	lg = logger
}
