package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"logging"
	"github.com/gin-gonic/gin"
	"commands"
)
// Global variables
var infolog *logrus.Logger

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
	infolog = logging.GetLogger ("INFO")
	infolog.WithFields(logrus.Fields{
		"package" : "main",
		"source": "main.go",
		"func": "main ()",
	}).Info("Expenses is starting !!!")

	route := gin.Default()
	//route.GET("/testing", startPage)
	route.PUT("/newuser", newuser)
	route.Run(":8085")
}

func newuser(c *gin.Context) {
	var user commands.CMD_NewUser

	if c.BindJSON(&user) == nil {
		fmt.Println("====== Bind By JSON ======")
		fmt.Println(user.Name)
		fmt.Println(user.Surname)
	}



//	c.String(200, "Success")
	c.JSON (200, gin.H {"message":"OK", "Cevap": 2})
}
