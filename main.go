package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	dir  string
	addr string
)

func main() {
	flag.StringVar(&dir, "dir", ".", "directory to serve")
	flag.StringVar(&addr, "addr", ":8000", "addr to serve files on")

	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(dir)))
	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
