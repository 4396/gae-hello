package hello

import (
	"net/http"
	"time"
)

func init() {
	http.HandleFunc("/hello", helloHandler)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello " + time.Now().String()))
}
