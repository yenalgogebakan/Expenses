package main

import (
	"fmt"
//	"github.com/sirupsen/logrus"
//	"logging"
	"github.com/gin-gonic/gin"
	"commands"
)

func newuser (c *gin.Context) {
	var user commands.CMD_NewUserParams

	if c.BindJSON(&user) == nil {
		fmt.Println (user.Name, user.Surname)
	}

	var cmdnewuser commands.CMD_NewUser
	cmdnewuser.P = &user

	err := cmdnewuser.Execute ()
	exception := ""
	if err != nil {
		exception = err.Error ()
	}
//	c.JSON (200, gin.H {"Name":user.Name, "Surname":user.Surname})
	c.JSON (200, gin.H {"exception":exception, "data":user})
}
