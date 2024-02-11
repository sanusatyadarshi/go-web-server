package http

import (
	"io"
	"log"
	"net/http"
)

func StartServer() {

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	signupHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Signup Here!\n")
	}
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/signup", signupHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
