# 07 Http

## Http Basics
### Http Handler 
Http handling is done using the `Handler` interface

### Http Handler interface
```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

### http.ServeMux
`ServeMux` is a simple multiplexer (mux). 
It implements the `Handler` interface and offers simple routing capabilities by mapping a given pattern to a http handler.
    
```go
func NewServeMux() *ServeMux
func (mux *ServeMux) Handle(pattern string, handler Handler)
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

### http.Server
`Server` is the actual http server being started.

```go
func (srv *Server) ListenAndServe() error
func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error
func (srv *Server) Serve(l net.Listener) error
```

### `DefaultServeMux`
For simple cases the default server and mux `DefaultServeMux` can be used.

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
})
http.ListenAndServe(":8080", nil)
```

### `http.Request`
```go
type Request struct {
        Method string
        URL *url.URL
        Header Header

        Body io.ReadCloser
        ContentLength int64
        Host string
        Form url.Values
        PostForm url.Values
        MultipartForm *multipart.Form
        ...
}

func (r *Request) Cookies() []*Cookie
func (r *Request) FormValue(key string) string
func (r *Request) ParseForm() error
func (r *Request) Referer() string
func (r *Request) UserAgent() string
...
```

### `http.ResponseWriter`
```go
type ResponseWriter interface {
        Header() Header
        Write([]byte) (int, error)
        WriteHeader(int)
}
```

## Handy features from the `http`-Package


### `http.FileServer`
Directory listing handler.

```go
func FileServer(root FileSystem) Handler
```

Example:
```go
http.Handle("/", http.FileServer(http.Dir(".")))
http.ListenAndServe(":8080", nil)
```

The `FileSystem` is a simple abstraction.
By implementing it other data sources can easily be passed to it.

## `http.Client`
Doing requests is just as simple as handling them.

```go
type Client
       func (c *Client) Do(req *Request) (resp *Response, err error)
       func (c *Client) Get(url string) (resp *Response, err error)
       func (c *Client) Head(url string) (resp *Response, err error)
       func (c *Client) Post(url string, bodyType string, body io.Reader) (resp *Response, err error)
       func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error)
```


```go
type Response struct {
        Status     string // e.g. "200 OK"
        StatusCode int    // e.g. 200
        Header Header
        Body io.ReadCloser
        ContentLength int64
        ...
}
```
## Exercise: File Storage Server
Create a simple Server, capable of serving and persisting files.

```shell
# Saving a file
# Antwort: Http 201
curl -X POST --data "some-text-data" 127.0.0.1:8080/path/to/file

# Serving a file
# Antwort: Http 200 "some-text-data"
curl 127.0.0.1:8080/path/to/file

# Error if file does not exist
# Antwort: Http 404
curl 127.0.0.1:8080/wrong/path
```

## Exercise: KV-store with REST interface
Expand the key value store with an interface.

```shell
# Saving a value
# Answer: Http 201
curl -X POST --data "some-text-data" 127.0.0.1:8080/key

# Getting a value
# Answer: Http 200 "some-text-data"
curl 127.0.0.1:8080/key

# Error if key does not exist
# Answer: Http 404
curl 127.0.0.1:8080/wrongKey
```

