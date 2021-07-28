package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "8443", "port to serve on")
	directory := flag.String("d", "/home/prateek/dev_local/go/src/Basic_Restaurant_Menu/Menu/", "menu")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	err := http.ListenAndServeTLS(":"+*port, "/home/prateek/server.crt", "/home/prateek/server.key", nil)

	if err != nil {
		log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
		log.Fatal("ListenAndServeTLS: ", err)
	}
}
