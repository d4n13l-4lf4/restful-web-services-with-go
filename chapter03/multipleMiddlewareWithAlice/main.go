package main

import (
	"github.com/justinas/alice"
	"log"
	"net/http"
	"strconv"
	"time"
)

func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currently in the check content type middleware")
		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type. Please send JSON"))
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r * http.Request) {
		handler.ServeHTTP(w, r)
		cookie := http.Cookie{Name: "Server-Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &cookie)
		log.Println("Currently in the set server time middleware")
	})
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("OK!"))
}

func main() {
	originalHandler := http.HandlerFunc(handle)
	chain := alice.New(filterContentType, setServerTimeCookie).Then(originalHandler)
	http.Handle("/city", chain)
	http.ListenAndServe(":8000", nil)
}