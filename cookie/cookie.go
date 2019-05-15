package cookie

import (
	"net/http"

	"github.com/garyburd/redigo/redis"
)

// CookieFunc is the function definition of how a function for cookies should
// look like. It requires the Responsewriter, Request, Name of the Cookie and
// the connection to the Redis pool.
type CookieFunc func(http.ResponseWriter, *http.Request, string, *redis.Pool) (Cookie, error)

// Cookie is the interface to define the cookie used by store values.
type Cookie interface {
	GetCookie() http.Cookie
	//GetConn() redis.PubSubConn
	GetSessionID() string
	SetSessionID(http.ResponseWriter, string)
	GetValue(string) []byte
	SetValue(string, interface{})
	GetUint64(string) uint64
	GetInt64(string) int64
	GetString(string) string
	GetBool(string) bool
	SetInterface(string, interface{}) error
	GetInterface(string, interface{}) error
	GetUint64Array(string) []uint64
	DeleteValue(string) error
	Remove(http.ResponseWriter)
}
