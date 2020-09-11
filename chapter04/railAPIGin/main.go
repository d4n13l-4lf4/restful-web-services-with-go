package main

import (
	"database/sql"
	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter04/dbutils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type StationResource struct {
	ID int `json:"id"`
	Name string `json:"name"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}

func GetStation(c *gin.Context) {
	var station StationResource
	id := c.Param("station_id")
	err := DB.QueryRow(dbutils.SELECT_STATION_BY_ID, id).Scan(&station.ID, &station.Name, &station.OpeningTime, &station.ClosingTime)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": station,
		})
	}
}

func CreateStation(c *gin.Context) {
	var station StationResource

	if err := c.BindJSON(&station); err == nil {
		statement, _ := DB.Prepare(dbutils.INSERT_STATION)
		result, _ := statement.Exec(station.Name, station.OpeningTime, station.ClosingTime)
		if err == nil {
			newID, _ := result.LastInsertId()
			station.ID = int(newID)
			c.JSON(http.StatusOK, gin.H{
				"result": station,
			})
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func RemoveStation(c *gin.Context) {
	id := c.Param("station-id")
	statement, _ := DB.Prepare(dbutils.DELETE_STATION_BY_ID)
	_, err := statement.Exec(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.String(http.StatusOK, "")
	}
}

func main() {
	var err error
	DB, err = sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed")
	}

	dbutils.Initialize(DB)
	r := gin.Default()
	r.GET("/v1/stations/:station_id", GetStation)
	r.POST("/v1/stations", CreateStation)
	r.DELETE("/v1/stations/:station_id", RemoveStation)
	r.Run(":8000")
}