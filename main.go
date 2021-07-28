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

	err := http.ListenAndServeTLS(":8443", "/home/prateek/dev_local/go/certs/pch-csr.cer", "/home/prateek/dev_local/go/certs/pch-pkey.pem", nil)

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)

	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}
