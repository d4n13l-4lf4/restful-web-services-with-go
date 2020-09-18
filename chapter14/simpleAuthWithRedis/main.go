package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gopkg.in/boj/redistore.v1"
)

var store *redistore.RediStore
var users = map[string]string{"daniel": "1234", "admin": "password"}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	if originalPassword, ok := users[username]; ok {
		session, _ := store.Get(r, "session.id")
		if password == originalPassword {
			session.Values["authenticated"] = true
			session.Save(r, w)
		} else {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "User is not found", http.StatusNotFound)
		return
	}
	w.Write([]byte("Logged In successfully"))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	session.Options.MaxAge = -1
	session.Save(r, w)
	w.Write([]byte(""))
}

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	if (session.Values["authenticated"] != nil) && session.Values["authenticated"] != false {
		w.Write([]byte(time.Now().String()))
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	store, err = redistore.NewRediStore(10, "tcp", ":6379", "", []byte(os.Getenv("SESSION_SECRET")))

	r := mux.NewRouter()
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/healthcheck", HealthcheckHandler)
	r.HandleFunc("/logout", LogoutHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr: "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}