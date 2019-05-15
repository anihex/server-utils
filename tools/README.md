# Tools

This part contains general purpose tools that help with common tasks such as
bind a JSON Object into an Object.

## ErrIfTrue

Generates an error, if the condition is true. This is usefull in combination with
the Error-Views of the views part.

```go
func handler(w http.ResponseWriter, r *http.Request) {
    err := tools.ErrIfTrue(r.Method != "GET", "Only GET Requests valid")
    if views.BadRequestIfErr(w, r, err) {
        return
    }

    // Continue the request ...
}
```

As opposed to this

```go
func handler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        views.SetError(r, "Only GET Requests are valid")
        views.ErrBadRequest(w, r)
        return
    }

    // Continue the request ...
}
```

## GID

A random string with the length of n. It can contain up to 64 different
characters.

## BindJSON

Takes a HTTP-Request and an interface. It reads the body from the Request and binds it into the interface.
If there was an error, it will be returned.

## Stringify

Takes an interface and returns a string that is encoded as JSON.

## GCD

Takes two int64 and calculates the "Greatest Common Divider" and returns it as an int64.

## MatMult

Takes two metrices and their width. It then calculates the multiplication of both and returns the new matrix.

## Ratio

Takes the Width, the Height and a Transformationmatrix. It calculates the ratio based on this.

## Determinant

Calcs the determinant of a 3x3 Matrix

## InverseAffine

Creates the inverse of affine transformation metrices.

## GetIP

Determines the IP from a request. It takes the built-in IP value and
additional headers from the request.

## NewRSA

Generates a new RSA Key-Pair with a name.

## RSA.Sign

Signs a message using the private key of the RSA-Pair.

## RSA.Verify

Checks the integrity of a message using the public key of the RSA-Pair.

## RSA.ExportPrivateKey

Exports the private key of a RSA-Pair into a file.

## RSA.ExportPublicKey

Exports the public key of a RSA-Pair into a file.

## RSA.ImportPrivateKey

Imports the private key of a RSA-Pair into a file.

## RSA.ImportPublicKey

Imports the public key of a RSA-Pair into a file.

## NewZMQ

Creates an instance of the ZMQ struct. It only contains the ZMQ4 socket

## ZMQ.SendDirect

It takes a URL, a ressource, a method and an object. It creates a new struct containing the ressource, the object, the method and the current time.

It then sends it to all ZMQ