// +build redis

package cookie

import (
	"net/http"

	"github.com/garyburd/redigo/redis"
)

// DummyCookie returns the http Cookie of the cookie
type DummyCookie struct {
	Values    map[string]interface{}
	SessionID string
}

// GetCookie returns a fresh http cookie
func (d *DummyCookie) GetCookie() http.Cookie {
	return http.Cookie{}
}

// GetSessionID returns the SessionID of the cookie
func (d *DummyCookie) GetSessionID() string {
	return d.SessionID
}

// SetSessionID forces a change of the session-ID
func (d *DummyCookie) SetSessionID(w http.ResponseWriter, id string) {
	d.SessionID = id
}

// GetValue reads a value from redis and returns it
func (d *DummyCookie) GetValue(Name string) []byte {
	return []byte{}
}

// SetValue stores a value in redis
func (d *DummyCookie) SetValue(Name string, value interface{}) {
	d.Values[Name] = value
}

// GetUint64 reads a value from redis and returs it as uint64
func (d *DummyCookie) GetUint64(Name string) uint64 {
	result, _ := d.Values[Name].(uint64)
	return result
}

// GetBool reads a value from redis and returns it as bool
func (d *DummyCookie) GetBool(Name string) bool {
	result, _ := d.Values[Name].(bool)
	return result
}

// GetInt64 reads a value from redis and returns it as int64
func (d *DummyCookie) GetInt64(Name string) int64 {
	result, _ := d.Values[Name].(int64)
	return result
}

// GetString reads a value from redis and returns it as string
func (d *DummyCookie) GetString(Name string) string {
	result, _ := d.Values[Name].(string)
	return result
}

// SetInterface stores an Interface using JSON encoding
func (d *DummyCookie) SetInterface(Name string, Value interface{}) error {
	d.Values[Name] = Value
	return nil
}

// GetInterface reads a JSON value from redis and binds it to the o interface
func (d *DummyCookie) GetInterface(Name string, o interface{}) error {
	o = d.Values[Name]
	return nil
}

// GetUint64Array reads a value from redis and returns it as uint64 array
func (d *DummyCookie) GetUint64Array(Name string) []uint64 {
	result, _ := d.Values[Name].([]uint64)
	return result
}

// DeleteValue deletes a value from redis
func (d *DummyCookie) DeleteValue(Name string) error {
	return nil
}

// Remove deletes all entries in redis. It also invalidates the http cookie
func (d *DummyCookie) Remove(w http.ResponseWriter) {}

// NewDummyCookie creates a new dummy cookie
func NewDummyCookie(Values map[string]interface{}, SessionID string) CookieFunc {
	return func(w http.ResponseWriter, r *http.Request, Name string, Conn *redis.Pool) (Cookie, error) {
		return &DummyCookie{
			Values:    Values,
			SessionID: SessionID,
		}, nil
	}
}
