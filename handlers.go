package main

import (
	"fmt"
//	"net/http"
//	"io/ioutil"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"logging"
	"github.com/gin-gonic/gin"
	"commands"
	"domain"
)

func servicemux (c *gin.Context) {
	services := make(map[string]interface{})
	var reqdata []byte
	var err error

	if reqdata, err = c.GetRawData (); err != nil {
		fmt.Println ("Cannot get reqdata")
		c.JSON (200, gin.H{"exception":"reqdata alinamadi"})
	}
	
	if err = json.Unmarshal(reqdata, &services); err != nil {
		fmt.Println ("error in json")
		fmt.Println (c.GetRawData())
	} else {
		fmt.Println ("switch e firdi")
		fmt.Println (services["servicename"])
		switch services ["servicename"] {
			case "GetuserById":
				fmt.Println ("Sername : GetUserById")
			case "ListUsers":
				fmt.Println ("Sername : ListUsers")
			case "GetExpenseById":
				fmt.Println ("Sername : GetEspenceById")
		}
	}
}

func newuser (c *gin.Context) {
	var user domain.User

	if c.BindJSON(&user) == nil {
		var cmdnewuser commands.CMD_NewUser
		cmdnewuser.P = &user

		err := cmdnewuser.Execute ()
		if err != nil {
			exception := err.Error ()
			logging.GetLogger("DEFAULT").WithFields(logrus.Fields{
			"package": "main",
			"source": "handlers.go",
			"func": "newuser",
			}).Error("commands.CMD_NewUser.Execute error")
			c.JSON (200, gin.H {"exception":exception, "data":user})
		}
	} else {
                logging.GetLogger("DEFAULT").WithFields(logrus.Fields{
                "package": "main",
                "source": "handlers.go",
                "func": "newuser",
                }).Error("can not decode JSON to user")
		exception := "CAN NOT DECEODE JSON TO USER STRUCT"
		c.JSON (200, gin.H {"exception":exception, "data":user})
        }

	c.JSON (200, gin.H {"exception":"SUCCESS", "data":user})
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
			}).Error("commands.CMD_NewExpense.Execute error")
			c.JSON (200, gin.H {"exception":exception, "data":expense})
		}
	} else {
                logging.GetLogger("DEFAULT").WithFields(logrus.Fields{
                "package": "main",
                "source": "handlers.go",
                "func": "newexpense",
                }).Error("can not decode JSON to Expense")
		exception := "CAN NOT DECEODE JSON TO EXPENCE STRUCT"
		c.JSON (200, gin.H {"exception":exception, "data":expense})
        }

	c.JSON (200, gin.H {"exception":"SUCCESS", "data":expense})
}
func getuser (c *gin.Context) {
	var user domain.User
	uname := c.Param("name")

	var cmdgetuser commands.CMD_GetUser
	cmdgetuser.P = &user
	cmdgetuser.Searchname = uname

	if err := cmdgetuser.Execute ();err != nil {
		exception := err.Error ()
		logging.GetLogger("DEFAULT").WithFields(logrus.Fields{
		"package": "main",
		"source": "handlers.go",
		"func": "getuser",
		}).Error("commands.CMD_GetUser.Execute error")
		c.JSON (200, gin.H {"exception":exception})
		return
	}
	c.JSON (200, gin.H {"exception":"SUCCESS", "data":cmdgetuser.P})
}
