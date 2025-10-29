package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := flag.String("port", "8000", "port to listen on")
	publicDir := "public"
	flag.Parse()
	fmt.Printf("Serving Go by Example at http://127.0.0.1:%s\n", *port)
	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(publicDir)))
}
