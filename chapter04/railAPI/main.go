package main

import (
	"database/sql"
	"encoding/json"
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
	"time"

	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter04/dbutils"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type TrainResource struct {
	ID              int
	DriverName      string
	OperatingStatus bool
}

type StationResource struct {
	ID int
	Name string
	OpeningTime time.Time
	ClosingTime time.Time
}

type ScheduleResource struct {
	ID int
	TrainID int
	StationID int
	ArrivalTime time.Time
}

func (t *TrainResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/v1/trains").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("{train-id}").To(t.removeTrain))
	container.Add(ws)
}

func (t TrainResource) getTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	err := DB.QueryRow(dbutils.SELECT_TRAIN_BY_ID, id).Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		handleError(err, "Train could not be found", http.StatusNotFound, response)
	} else {
		response.WriteEntity(t)
	}
}

func (t TrainResource) createTrain(request *restful.Request, response *restful.Response) {
	log.Println(request.Request.Body)
	decoder := json.NewDecoder(request.Request.Body)
	var b TrainResource
	err := decoder.Decode(&b)
	log.Println(b.DriverName, b.OperatingStatus)
	if err != nil {
		handleError(err, "Cannot decode request body", http.StatusBadRequest, response)
		return
	}

	statement, err := DB.Prepare(dbutils.INSERT_TRAIN)
	if err != nil {
		handleError(err, err.Error(), http.StatusInternalServerError, response)
		return
	}

	result, err := statement.Exec(b.DriverName, b.OperatingStatus)
	if err == nil {
		newID, _ := result.LastInsertId()
		b.ID = int(newID)
		response.WriteHeaderAndEntity(http.StatusCreated, b)
	} else {
		handleError(err, err.Error(), http.StatusInternalServerError, response)
	}

}

func (t TrainResource) removeTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	statement, _ := DB.Prepare(dbutils.DELETE_TRAIN_BY_ID)
	_, err := statement.Exec(id)
	if err == nil {
		response.WriteHeader(http.StatusOK)
	} else {
		handleError(err, err.Error(), http.StatusInternalServerError, response)
	}
}

func handleError(err error, message string, status int, response *restful.Response) {
	log.Println(err)
	response.AddHeader("Content-Type", "text/plain")
	response.WriteErrorString(status, message)
}

func main() {
	var err error
	DB, err = sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}

	dbutils.Initialize(DB)
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	t := TrainResource{}
	t.Register(wsContainer)
	log.Printf("start listening on localhost:8000")
	server := &http.Server{Addr: ":8000", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}