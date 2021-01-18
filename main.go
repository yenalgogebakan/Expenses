package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"logging"
	"repository"
)
// Global variables
var infolog *logrus.Logger
var route *gin.Engine

func main () {
	// Log setup
	if err := logging.CreateDefaultLogs (); err != nil {
		fmt.Println ("Can not initialize logs, exiting")
		return
	}
	infolog = logging.GetLogger ("INFO")
	infolog.WithFields(logrus.Fields{
		"package" : "main",
		"source": "main.go",
		"func": "main ()",
	}).Info("Expenses is starting !!!")

	route = gin.Default()
	prepareHandlers ()
	prepareRepository ()
	route.Run(":8085")

}

func prepareHandlers () {
// Simple group: v1
	v1 := route.Group("/v1")

	v1.PUT("/users", newuser)
	v1.PUT("/expenses", newexpense)
	v1.GET("/users/:name", getuser)
	v1.POST("/services", servicemux)
//
//	v1.PUT("/expenses", newexpense)
//	v1.GET("/expenses", getexpense)
//	v1.POST("/expenses", proexpense)
}

func prepareRepository () {
	db, err := sql.Open("sqlite3", "Expenses.db")
	if err != nil { panic(err) }
	if db == nil { panic("db nil") }
	repository.Mainrepo = repository.NewSqliteRepository (db)
/*
//Inmem
	inmemdbpointer, err := inmempoool ("Expense")
	if err != nill {panic (err)}
	if inmemdbpointer == nil {panic(inmemdbpointer nil)}
	repository.Mainrepo = repository.NewInmemRepository (inmemdbpointer)
*/
}
