### HandleFunc

HandleFunc registers the handler function for the given pattern in the DefaultServeMux.   
The documentation for ServeMux explains how patterns are matched.

```
HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

### ResponseWriter parameter

// A ResponseWriter interface is used by an HTTP handler to  
// construct an HTTP response.  
 
// A ResponseWriter may not be used after the Handler.ServeHTTP method  
// has returned.  
type ResponseWriter interface {}

### Request parameter

// A Request represents an HTTP request received by a server  
// or to be sent by a client.  

// The field semantics differ slightly between client and server  
// usage. In addition to the notes on the fields below, see the  
// documentation for Request.Write and RoundTripper.  
type Request struct {}
