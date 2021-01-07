package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"logging"
	"github.com/gin-gonic/gin"
)
// Global variables
var mainlog, infolog *logrus.Logger

type Person struct {
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
}
func main () {
	// Log setup
	if err := logging.CreateDefaultLogs (); err != nil {
		fmt.Println ("Can not initialize logs, exiting")
		return
	}
	if tmplog, err := logging.GetLogger ("DEFAULT"); err != nil {
		fmt.Println ("main DEFAULT log can not be accessed, exiting")
		return
	} else {
		mainlog = tmplog
	}
	if tmplog, err := logging.GetLogger ("INFO"); err != nil {
		fmt.Println ("main INFO log can not be accessed, exiting")
		return
	} else {
		infolog = tmplog
	}
	infolog.WithFields(logrus.Fields{
		"package" : "main",
		"source": "main.go",
		"func": "main ()",
	}).Info("Expenses is starting !!!")

	route := gin.Default()
	route.GET("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	if c.Bind(&person) == nil {
		fmt.Println("====== Bind By Query String ======")
		fmt.Println(person.Name)
		fmt.Println(person.Address)
	}
	if c.BindJSON(&person) == nil {
		fmt.Println("====== Bind By JSON ======")
		fmt.Println(person.Name)
		fmt.Println(person.Address)
	}

//	c.String(200, "Success")
c.JSON (200, gin.H {"message":"OK", "Cevap": 2})
}
