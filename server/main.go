package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count int64
var m = &sync.Mutex{}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./assets/index.html")
}

func incHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprint(w, count)
	case "POST":
		m.Lock()
		count++
		m.Unlock()
	default:
		fmt.Fprint(w, "Only GET and POST requests are supported.")
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(`./assets`)))
	http.HandleFunc("/inc", incHandler)

	fmt.Println("Server listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
