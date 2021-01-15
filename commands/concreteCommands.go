package commands

import (
        "domain"
//	"fmt"
//        "errors"
        "logging"
        "github.com/sirupsen/logrus"
	"math/rand"
	"time"
	"repository"
)

// New user Command
type CMD_NewUser struct {
	P *domain.User
}
func (cmd_newuser CMD_NewUser) Execute () error {
	rand.Seed(time.Now().UnixNano())
	cmd_newuser.P.UserId = rand.Intn (1000)

	_, err := repository.Mainrepo.StoreUser (cmd_newuser.P)
	if err != nil {
		logging.GetLogger("DEFAULT").WithFields(logrus.Fields{
		"package":"commands",
		"source":"concreteCommands.go",
		"func":"CMD_NewUser",
		"username":cmd_newuser.P.Name,
		}).Error("repository.Mainrepo.StoreUser error")
		return err
	}

	logging.GetLogger("INFO").WithFields(logrus.Fields{
        "package": "commands",
        "source":"concreteCommands.go",
        "func": "CMD NewUser",
        }).Info("command executed successfully")
	return nil
}

// New expense Command
type CMD_NewExpense struct {
	P *domain.Expense
}
func (cmd_newexpense CMD_NewExpense) Execute () error {
	rand.Seed(time.Now().UnixNano())
	cmd_newexpense.P.ExpenseId = rand.Intn (1000)

	_, err := repository.Mainrepo.StoreExpense (cmd_newexpense.P)
	if err != nil {
		logging.GetLogger("DEFAULT").WithFields(logrus.Fields{
		"package":"commands",
		"source":"concreteCommands.go",
		"func":"CMD_NewExpense",
		"expitem":cmd_newexpense.P.ExpItem,
		}).Error("repository.Mainrepo.StoreExpense error")
		return err
	}

	logging.GetLogger("INFO").WithFields(logrus.Fields{
        "package": "commands",
        "source":"concreteCommands.go",
        "func": "CMD NewExpense",
        }).Info("command executed successfully")
	return nil
}

// Get user Command
type CMD_GetUser struct {
	P *domain.User
	Searchname string
}
func (cmd_getuser CMD_GetUser) Execute () error {
	var err error
	var tmpuser *domain.User

	tmpuser, err = repository.Mainrepo.GetUserByName (cmd_getuser.Searchname)
	if err != nil {
		logging.GetLogger("DEFAULT").WithFields(logrus.Fields{
		"package":"commands",
		"source":"concreteCommands.go",
		"func":"CMD_GetUser",
		"username":cmd_getuser.Searchname,
		}).Error("repository.Mainrepo.StoreGetUserByName error")
		return err
	}
	*(cmd_getuser.P) = *(tmpuser)
	logging.GetLogger("INFO").WithFields(logrus.Fields{
        "package": "commands",
        "source":"concreteCommands.go",
        "func": "CMD GetUser",
        }).Info("command executed successfully")
	return nil
}

