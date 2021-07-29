package main

import (
	// "fmt"
	// "io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func main() {
	log.Println("1.Starting TLS on: 8443")
	http.HandleFunc("/", HelloServer)
	log.Println("2.Starting TLS on: 8443")
	err := http.ListenAndServeTLS(":8443", "/home/prateek/pch.crt", "/home/prateek/pch-pkey.pem", nil)
	log.Println("Starting TLS on: 8443")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
