package main

import (
	"database/sql"
	"encoding/json"
	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter07/base62Example/base62"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter07/urlshortener/helper"
	"github.com/gorilla/mux"
	)

type DBClient struct {
	db *sql.DB
}

type Record struct {
	ID int `json:"id"`
	URL string `json:"url"`
}

func (driver *DBClient) GenerateShortURL(w http.ResponseWriter, r *http.Request) {
	var id int
	var record Record
	postBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(postBody, &record)
	err = driver.db.QueryRow(helper.INSERT_WEB_URL, record.URL).Scan(&id)
	responseMap := map[string]string{"encoded_string": base62.ToBase62(id)}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

func (driver *DBClient) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	var url string
	vars := mux.Vars(r)
	id := base62.ToBase10(vars["encoded_string"])
	err := driver.db.QueryRow(helper.SELECT_WEB_URL_BY_ID, id).Scan(&url)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		responseMap := map[string]interface{}{"url": url}
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

func main() {

	db, err := helper.InitDB()
	if err != nil {
		panic(err)
	}

	dbclient := &DBClient{db: db}
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/v1/short/{encoded_string:[a-zA-Z0-9]*}", dbclient.GetOriginalURL).Methods("GET")
	r.HandleFunc("/v1/short", dbclient.GenerateShortURL).Methods("POST")
	srv := &http.Server{
		Handler: r,
		Addr: "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}