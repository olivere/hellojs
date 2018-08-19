# HelloJS

This repository represents an example of using WebAssembly with
Go 1.11.

You need [Caddy](https://caddyserver.com/) or any other server to
host the web application.

You also need to [enable WebAssembly in Chrome](https://www.nextofwindows.com/how-to-enable-webassembly-in-chrome).

Next, compile the WebAssembly module:

```
$ make
GOOS=js GOARCH=wasm go1.11rc1 build -o app/main.wasm
```

P.S.: You might need to change the Go executable once Go 1.11 final
is out. I've use `go1.11rc1` in the example above.

Then run the HTTP server with Caddy and open the web page at
[http://localhost:8080](http://localhost:8080):

```
$ caddy
Activating privacy features... done.
http://localhost:8080
$ open http://localhost:8080
```

# License

MIT
