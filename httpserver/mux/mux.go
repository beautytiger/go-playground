package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/", apiHandler{})
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprint(w, "welcome to index")
	})
	var srv = &http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	log.Fatal(srv.ListenAndServe())
}

type apiHandler struct{}

func (a apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "api")
}
