package cookie

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/anihex/server-utils/tools"
	"github.com/garyburd/redigo/redis"
)

// RedisCookie ist ein einfaches Interface für Redis basierte Sessions
type RedisCookie struct {
	Cookie    http.Cookie
	Pool      *redis.Pool
	SessionID string
	stored    bool
	w         http.ResponseWriter
	r         *http.Request
	name      string
}

// GetCookie returns the http Cookie of the cookie
func (session *RedisCookie) GetCookie() http.Cookie {
	return session.Cookie
}

// GetConn returns the redis pool of the cookie
func (session *RedisCookie) GetConn() *redis.Pool {
	return session.Pool
}

// GetSessionID returns the SessionID of the cookie
func (session *RedisCookie) GetSessionID() string {
	return session.SessionID
}

// GetValue reads a value from redis and returns it
func (session *RedisCookie) GetValue(Name string) []byte {
	resultObj, err := session.Pool.Get().Do("HGET", session.SessionID, Name)
	if err != nil {
		return []byte{}
	}

	resultByte, _ := resultObj.([]byte)

	return resultByte
}

// SetValue stores a value in redis
func (session *RedisCookie) SetValue(Name string, Value interface{}) {
	session.Pool.Get().Do("HSET", session.SessionID, Name, Value)

	session.Store()
}

// GetUint64 reads a value from redis and returs it as uint64
func (session *RedisCookie) GetUint64(Name string) uint64 {
	data := session.GetValue(Name)
	var result uint64
	err := json.Unmarshal(data, &result)

	if err != nil {
		result = 0
	}

	return result
}

// GetBool reads a value from redis and returns it as bool
func (session *RedisCookie) GetBool(Name string) bool {
	data := session.GetValue(Name)

	var result bool
	err := json.Unmarshal(data, &result)

	log.Println(data, result)

	if err != nil {
		result = false
	}

	return result
}

// GetInt64 reads a value from redis and returns it as int64
func (session *RedisCookie) GetInt64(Name string) int64 {
	data := session.GetValue(Name)
	var result int64
	err := json.Unmarshal(data, &result)

	if err != nil {
		result = 0
	}

	return result
}

// GetString reads a value from redis and returs it as string
func (session *RedisCookie) GetString(Name string) string {
	value := session.GetValue(Name)

	return fmt.Sprintf("%s", value)
}

// SetInterface stores an Interface using JSON encoding
func (session *RedisCookie) SetInterface(Name string, Value interface{}) error {
	ToStore, err := json.Marshal(Value)
	if err != nil {
		return err
	}

	session.SetValue(Name, ToStore)

	return nil
}

// GetInterface reads a JSON value from redis and binds it to the o interface
func (session *RedisCookie) GetInterface(Name string, o interface{}) error {
	data := session.GetValue(Name)
	err := json.Unmarshal(data, &o)

	return err
}

// GetUint64Array reads a value from redis and returns it as uint64 array
func (session *RedisCookie) GetUint64Array(Name string) []uint64 {
	data := session.GetValue(Name)
	var result []uint64
	json.Unmarshal(data, &result)

	return result
}

// DeleteValue deletes a value from redis
func (session *RedisCookie) DeleteValue(Name string) error {
	session.Pool.Get().Do("HDEL", session.SessionID, Name)

	session.Store()

	return nil
}

// Remove deletes all entries in redis. It also invalidates the http cookie
func (session *RedisCookie) Remove(w http.ResponseWriter) {
	session.Pool.Get().Do("DEL", session.SessionID)
	session.SessionID = ""
	session.Cookie.Value = ""
	session.Cookie.Expires = time.Unix(0, 0)

	http.SetCookie(w, &session.Cookie)
}

// SetSessionID forces a change of the session-ID
func (session *RedisCookie) SetSessionID(w http.ResponseWriter, id string) {
	session.Cookie.Value = id
	session.SessionID = id

	http.SetCookie(w, &session.Cookie)
}

// Store saves the http-Cookie if neccessary
func (session *RedisCookie) Store() {
	if session.stored {
		return
	}

	if _, err := session.r.Cookie(session.name); err != nil {
		http.SetCookie(session.w, &session.Cookie)
	}

	session.stored = true
}

// NewRedisCookie creates a new redis cookie
func NewRedisCookie(w http.ResponseWriter, r *http.Request, Name string, Conn *redis.Pool) (Cookie, error) {
	var token string

	cookie, err := r.Cookie(Name)
	if err != nil {
		token = tools.GID(32)

		if w == nil {
			return &RedisCookie{}, errors.New("responseWriter not set")
		}

		if r == nil {
			return &RedisCookie{}, errors.New("request not set")
		}

		c := http.Cookie{
			Name:    Name,
			Path:    "/",
			Expires: time.Unix(time.Now().Unix()+(60*60*24*365*10), 0),
			Value:   token,
			//Secure:   true,
			HttpOnly: true,
		}

		return &RedisCookie{
			Pool:      Conn,
			SessionID: c.Value,
			Cookie:    c,
			w:         w,
			r:         r,
			name:      Name,
		}, nil
	}

	if strings.TrimSpace(cookie.Value) == "" {
		return &RedisCookie{}, errors.New("request contained empty value")
	}

	if cookie.Value != "0" {
		token = cookie.Value
	}

	cookie.Path = "/"
	result := &RedisCookie{
		Pool:      Conn,
		SessionID: token,
		Cookie:    *cookie,
		w:         w,
		r:         r,
		name:      Name,
	}

	return result, nil
}
