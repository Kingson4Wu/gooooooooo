+ https://gowebexamples.com/routes-using-gorilla-mux/

+ The biggest strength of the gorilla/mux Router is the ability to extract segments from the request URL. As an example, this is a URL in your application:

/books/go-programming-blueprint/page/10

+ Setting the HTTP server’s router
Ever wondered what the nil in http.ListenAndServe(":80", nil) ment? It is the parameter for the main router of the HTTP server. By default it’s nil, which means to use the default router of the net/http package. To make use of your own router, replace the nil with the variable of your router r.

http.ListenAndServe(":80", r)

+ Restrict the request handler to specific HTTP methods.

r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

+ Restrict the request handler to specific hostnames or subdomains.

r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

+ Restrict the request handler to http/https.

r.HandleFunc("/secure", SecureHandler).Schemes("https")
r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

+ Restrict the request handler to specific path prefixes.

bookrouter := r.PathPrefix("/books").Subrouter()
bookrouter.HandleFunc("/", AllBooks)
bookrouter.HandleFunc("/{title}", GetBook)



