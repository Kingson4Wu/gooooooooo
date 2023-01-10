+ https://gowebexamples.com/

+ https://gowebexamples.com/advanced-middleware/ !!!!! 很优雅的技巧

+ Middleware (Advanced)
This example will show how to create a more advanced version of middleware in Go.

A middleware in itself simply takes a http.HandlerFunc as one of its parameters, wraps it and returns a new http.HandlerFunc for the server to call.

Here we define a new type Middleware which makes it eventually easier to chain multiple middlewares together. This idea is inspired by Mat Ryers’ talk about Building APIs. You can find a more detailed explaination including the talk here.

This snippet explains in detail how a new middleware is created. In the full example below, we reduce this version by some boilerplate code.

+ https://gowebexamples.com/websockets/

