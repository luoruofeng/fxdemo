# fxdemo

```go
// NewHTTPServer builds an HTTP server that will begin serving requests
// when the Fx application starts.
func NewHTTPServer(lc fx.Lifecycle) *http.Server {
  srv := &http.Server{Addr: ":8080"}
  return srv
}
```


```go
func NewHTTPServer(lc fx.Lifecycle) *http.Server {
  srv := &http.Server{Addr: ":8080"}
  lc.Append(fx.Hook{
    OnStart: func(ctx context.Context) error {
      ln, err := net.Listen("tcp", srv.Addr)
      if err != nil {
        return err
      }
      fmt.Println("Starting HTTP server at", srv.Addr)
      go srv.Serve(ln)
      return nil
    },
    OnStop: func(ctx context.Context) error {
      return srv.Shutdown(ctx)
    },
  })
  return srv
}
 
```

 
```go
func main() {
  fx.New(
    fx.Provide(NewHTTPServer),
  ).Run()
}
```

```go
//Huh? Did something go wrong? The first line in the output states that the server was provided, but it doesn't include our "Starting HTTP server" message. The server didn't run.

  fx.New(
    fx.Provide(NewHTTPServer),
    fx.Invoke(func(*http.Server) {}),
  ).Run()
 
```


```shell
curl http://localhost:8080
```