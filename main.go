package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"logging"
	"github.com/gin-gonic/gin"
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
	defineHandlers ()
	route.Run(":8085")

}
/*
func newuser(c *gin.Context) {
	var user commands.CMD_NewUserParams

	if c.BindJSON(&user) == nil {
		fmt.Println("====== Bind By JSON ======")
		fmt.Println(user.Name)
		fmt.Println(user.Surname)
	}



//	c.String(200, "Success")
	c.JSON (200, gin.H {"Name":user.Name, "Surname": user.Surname})
}
*/
func defineHandlers () {
// Simple group: v1
	v1 := route.Group("/v1")

	v1.PUT("/users", newuser)
//	v1.GET("/users", getuser)
//
//	v1.PUT("/expenses", newexpense)
//	v1.GET("/expenses", getexpense)
//	v1.POST("/expenses", proexpense)
}
