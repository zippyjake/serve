package main

import (
	"log"
	"net/http"
	"os"
)

func serve(w http.ResponseWriter, r *http.Request) {
	file := "www" + r.URL.Path
	http.ServeFile(w, r, "www/"+r.URL.Path[1:])
	log.Print(file)
}

func main() {
	var port = ":8026"
	if len(os.Args) == 2 {
		port = ":" + os.Args[1]
	}
	http.HandleFunc("/", serve)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	log.Print("SERVING www on " + port)
}
