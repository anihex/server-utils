package middleware

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

var rd *redis.Pool
var lg *log.Logger

// SetRedis sets the package Redis Pool
func SetRedis(Pool *redis.Pool) {
	rd = Pool
}

// SetLog sets the package logger
func SetLog(logger *log.Logger) {
	lg = logger
}
