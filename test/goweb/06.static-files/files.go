package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}

/**
$ tree assets/
assets/
└── css
    └── styles.css
*/
/**
$ curl -s http://localhost:8080/static/css/styles.css
body {
    background-color: black;
}
*/
