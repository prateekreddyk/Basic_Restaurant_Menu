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

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)

	err := http.ListenAndServeTLS(":"+*port, "/home/prateek/pch.crt", "/home/prateek/pch.key", nil)

	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}

// package main

// import (
// 	// "fmt"
// 	// "io"
// 	"log"
// 	"net/http"
// )

// func HelloServer(w http.ResponseWriter, req *http.Request) {
// 	w.Header().Set("Content-Type", "text/plain")
// 	w.Write([]byte("This is an example server.\n"))
// 	// fmt.Fprintf(w, "This is an example server.\n")
// 	// io.WriteString(w, "This is an example server.\n")
// }

// func main() {
// 	log.Println("1.Starting TLS on: 8443")
// 	http.HandleFunc("/", HelloServer)
// 	log.Println("2.Starting TLS on: 8443")
// 	err := http.ListenAndServeTLS(":8443", "/home/prateek/pch.crt", "/home/prateek/pch-pkey.pem", nil)
// 	log.Println("Starting TLS on: 8443")
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }
