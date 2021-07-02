package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<!DOCTYPE html>
		<html>
			<head>
				<title>Hello, world</title>
			</head>	
			<body>
				Hello, world!
			</body>
		</html>`,
	)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
	// To serve static files
	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("assets")),
		),
	)
}
