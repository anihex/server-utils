# Middleware

Middleware that can be used in combination with [Vestigo](http://github.com/husobee/vestigo/).

The middleware consists of:

- CORS
- Dummy
- Log
- Time

## CORS

Adds CORS Informations to the Response Header.

```go
func main() {
    cors := middleware.CORS("https://github.com/, https://www.google.de/", 3600, true, "GET, OPTIONS")

    router := vestigo.NewRouter()
    router.Get("/", handler, cors)

    http.ListenAndServe(":8080", router)
}
```

## Dummy

The dummy can be used to replace existing middlewares. This can be usefull if
the middlewares are stored in a variable.

## Log

Logs incomming requests.

## Time

Adds the current time to the request context. This can be usefull to measure the
time it took to send a response.