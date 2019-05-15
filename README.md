# Server Utils

This package contains some of the utils I use for my go http servers.

## Containing tools:

- Cookie - A redis based cookie that uses a single token on the client side
- Middleware - Some HTTP Middleware that can be used with [Vestigo](https://github.com/husobee/vestigo) such as loggin etc.
- Tools - A misc collection of tools such as getting the IP from a request (forwarded by reverse proxy) etc.
- Views - Some simple functions to send JSON, error messages and files to the client
