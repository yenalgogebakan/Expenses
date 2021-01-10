package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"logging"
	"github.com/gin-gonic/gin"
	"commands"
	"domain"
)

func newuser (c *gin.Context) {
	var user domain.User

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
	c.JSON (200, gin.H {"exception":exception, "data":user})
}

func newexpense (c *gin.Context) {
	var expense domain.Expense

	if c.BindJSON(&expense) == nil {
		var cmdnewexpense commands.CMD_NewExpense
		cmdnewexpense.P = &expense

		err := cmdnewexpense.Execute ()
		if err != nil {
			exception := err.Error ()
                	logging.GetLogger("DEFAULT").WithFields(logrus.Fields{
                	"package": "main",
                	"source": "handlers.go",
                	"func": "newexpense",
                	}).Error("errtxt-error in new expense command")
			c.JSON (200, gin.H {"exception":exception, "data":expense})
		}
	} else {
                logging.GetLogger("DEFAULT").WithFields(logrus.Fields{
                "package": "main",
                "source": "handlers.go",
                "func": "newexpense",
                }).Error("errtxt-can not decode JSON to Expense")
		exception := "CAN NOT DECEODE JSON TO EXPENCE STRUCT"
		c.JSON (200, gin.H {"exception":exception, "data":expense})
        }

	c.JSON (200, gin.H {"exception":"SUCCESS", "data":expense})
}
