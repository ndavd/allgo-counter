package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

	var port string
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	} else {
		port = "3000"
	}
	fmt.Println("Server listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
