package controllers

import (
	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter04/dbutils"
	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter04/railAPIRevel/app"
	"github.com/revel/revel"
	"log"
	"net/http"
)

type App struct {
	*revel.Controller
}

type TrainResource struct {
	ID int `json:"id"`
	DriverName string `json:"driver_name"`
	OperatingStatus bool `json:"operating_status"`
}

func (c App) GetTrain() revel.Result {
	var train TrainResource
	id := c.Params.Route.Get("train-id")

	statement, _ := app.DB.Prepare(dbutils.SELECT_TRAIN_BY_ID)
	err := statement.QueryRow(id).Scan(&train.ID, &train.DriverName, &train.OperatingStatus)

	if err != nil {
		c.Response.Status = http.StatusInternalServerError
		return c.RenderText("Error my friend!")
	}

	c.Response.Status = http.StatusOK
	return c.RenderJSON(train)
}

func (c App) CreateTrain() revel.Result {
	var train TrainResource
	c.Params.BindJSON(&train)
	statement, _ := app.DB.Prepare(dbutils.INSERT_STATION)
	result, err := statement.Exec(train.DriverName, train.OperatingStatus)

	if err != nil {
		c.Response.Status = http.StatusInternalServerError
		return c.Render()
	}

	newID, _ := result.LastInsertId()
	train.ID = int(newID)
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(train)
}

func (c App) RemoveTrain() revel.Result {
	id := c.Params.Route.Get("train-id")
	statement, _ := app.DB.Prepare(dbutils.INSERT_TRAIN)
	_, err := statement.Exec(id)

	if err != nil {
		c.Response.Status = http.StatusInternalServerError
		return c.Render()
	}

	log.Println("Successfully deleted the resource")
	c.Response.Status = http.StatusOK
	return c.RenderText("")
}

func (c App) Index() revel.Result {
	return c.Render()
}
