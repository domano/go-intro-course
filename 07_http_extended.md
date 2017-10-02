# 08 http extended

## TLS & HTTP2 with Go
Sinve 1.6 go has support for HTTP2 out of the box using the existing http API.
As HTTP2 requires TLS, it only works if you use the TLS-function.

```go
panic(http.ListenAndServeTLS(":8443", "server.pem", "server.key", nil))
```
## Middleware

### Chaining
Handlers can easily be chained (like decorators).
Generic handlers can be used as middlewares, offering pluggable functionality.

```go
loggingMiddleware(delegate http.Handler) http.Handler
```

Example:
```go
	chain := LoggingMiddleware(
		AccessMiddleware(
			helloWorld
		)
	)
	panic(http.ListenAndServe(":8080", chain))
```


### Chaining with `alice`
https://github.com/justinas/alice

```go
	chain := alice.New(
		LoggingMiddleware,
		AccessMiddleware,
	).Then(helloWorld)
```

### Handlers of the Gorilla Toolkit
Handlers are generic and can be combined from varying frameworks and libraries.

http://www.gorillatoolkit.org/pkg/handlers


## Package `http/httptest`

#### ResponseRecorder
Easy testing of http handlers.

```go
handler := func(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "something failed", http.StatusInternalServerError)
}

req, err := http.NewRequest("GET", "http://example.com/foo", nil)
if err != nil {
	log.Fatal(err)
}

w := httptest.NewRecorder()
handler(w, req)

fmt.Printf("%d - %s", w.Code, w.Body.String())
```

#### httptest.NewServer
Starts a server on a random free port.

```go
ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, client")
}))
defer ts.Close()

res, err := http.Get(ts.URL)
```

## Web Frameworks & alternative Router

### `julienschmidt/httprouter`

Fast router with param parsing, using a search tree internally.

```go
import "github.com/julienschmidt/httprouter"


mux := httprouter.New()
mux.GET("/hello/:name", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "Hello %v\n", params.ByName("name"))
})
```

### Gorilla Toolkit
The gorilla toolkit offers a variety of packages, which can be used independently of each other.

http://www.gorillatoolkit.org

- `gorilla/context` stores global request variables.
- `gorilla/mux` is a powerful URL router and dispatcher.
- `gorilla/reverse` produces reversible regular expressions for regexp-based muxes.
- `gorilla/rpc` implements RPC over HTTP with codec for JSON-RPC.
- `gorilla/schema` converts form values to a struct.
- `gorilla/securecookie` encodes and decodes authenticated and optionally encrypted cookie values.
- `gorilla/sessions` saves cookie and filesystem sessions and allows custom session backends.
- `gorilla/websocket` implements the WebSocket protocol defined in RFC 6455.


## Übung: Testen eines Handlers
Write a handler using the gorilla/mux or julienschmidt/httprouter and verify it getting called by using the httptest Recorder.
