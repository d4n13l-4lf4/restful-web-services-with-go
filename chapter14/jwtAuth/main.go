package main

import (
	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter14/jwtAuth/auth"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var users = map[string]string {
	"daniel": "1234",
	"admin": "admin",
}

func HandleError(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	w.Write([]byte(message))
}

func GetToken(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		HandleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	username, password := r.PostForm.Get("username"), r.PostForm.Get("password")

	if originalPassword, ok := users[username] ; ok {
		if originalPassword == password {
			token, _ := auth.DeliverToken(map[string]interface{}{})
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(token))
		} else {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		}
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
	}
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(time.Now().String()))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/token", GetToken).Methods("POST")
	r.HandleFunc("/healthcheck", auth.Middleware(Healthcheck)).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr: "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}