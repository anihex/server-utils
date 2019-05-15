# Cookie

The cookie part allows the usage of the default HTTP Cookie, but stores the
values in redis.

```go
// Create cookie and store values
func PostHandler(w http.ResponseWriter, r *http.Request) {
    conn := GetRedisPool()

    c, err := cookie.NewRedisCookie(w, r, "mycookie", conn)
    if err != nil {
        log.Fatal(err)
    }

    c.SetValue("name", "demo")

    // Important! The http cookie will only be set AFTER this was called
    // AT LEAST once!
    c.Store()
}

// Afterwards, read the value from redis
func GetHandler(w http.ResponseWriter, r *http.Request) {
    conn := GetRedisPool()

    c, err := cookie.NewRedisCookie(w, r, "mycookie", conn)
    if err != nil {
        log.Fatal(err)
    }

    // Output: "demo"
    log.Printf("Name: %s", c.GetString("name"))
}
```
