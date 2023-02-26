# Handler Interface

### Type Handler
A Handler responds to an HTTP request.

ServeHTTP should write reply headers and data to the ResponseWriter and then return. Returning signals that the request is finished; it is not valid to use the ResponseWriter or read from the Request.Body after or concurrently with the completion of the ServeHTTP call.

Depending on the HTTP client software, HTTP protocol version, and any intermediaries between the client and the Go server, it may not be possible to read from the Request.Body after writing to the ResponseWriter. Cautious handlers should read the Request.Body first, and then reply.

Except for reading the body, handlers should not modify the provided Request.

If ServeHTTP panics, the server (the caller of ServeHTTP) assumes that the effect of the panic was isolated to the active request. It recovers the panic, logs a stack trace to the server error log, and either closes the network connection or sends an HTTP/2 RST_STREAM, depending on the HTTP protocol. To abort a handler so the client sees an interrupted response but the server doesn't log an error, panic with the value ErrAbortHandler.

type Handler interface {  
    ServeHTTP(ResponseWriter, *Request)  
}

以上关于 Handler接口 内容来自  
http://localhost:6060/pkg/net/http/#Handler


### HandleFunc 与 Handle 的区别
```
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))

Handle(pattern string, handler Handler)
```
首先在参数上，第二个参数有区别：Handle的 Handler 是一个接口，而 HandleFunc的 handler 是 Func  ------ Handler 实际调用就是 func(ResponseWriter, *Request) 函数类型


既然知道了 Handle的第二参数是个接口，那我们怎么模拟一个像 HandleFunc 一样的调用的呢？

1. 从接口入手  

声明一个 Hello type，并实现 Handler 实现，即可使用在 Handle 函数的第二参数上 
```
type Hello struct {

}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, " %d\n", r.URL. ...)
}
```
