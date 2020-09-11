package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got parameter id:%s!\n", queryParams["id"][0])
	fmt.Fprintf(w, "Got parameter category:%s!", queryParams["category"][0])
}

func main() {

	r := mux.NewRouter()
	// r.StrictSlash(true) /articles and /articles/ are the same
	// r.UseEncodedPath() /articles/2 is the same as /articles%2F2
	r.HandleFunc("/articles", QueryHandler)
	svr := http.Server{
		Handler: r,
		Addr: "localhost:8000",
		ReadTimeout: 15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Fatal(svr.ListenAndServe())

}