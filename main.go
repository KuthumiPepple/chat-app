package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
		<html>
		  <head>
		    <title>chat</title>
		  </head>
		  <body>
		    Let's chat!
		  </body>
		</html>`))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
