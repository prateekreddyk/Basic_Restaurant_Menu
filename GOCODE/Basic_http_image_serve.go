package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	//port := flag.String("p", "9000", "port to serve on")
	directory := flag.String("d", "/home/prateek/dev_local/go/src/Basic_Restaurant_Menu/Menu/", "menu")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	err := http.ListenAndServeTLS(":"+*port, "/home/prateek/dev_local/go/src/Basic_Restaurant_Menu/CERTS/pch-csr.cer", "/home/prateek/dev_local/go/src/Basic_Restaurant_Menu/CERTS/pch-pkey.pem", nil)

	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
		log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
}