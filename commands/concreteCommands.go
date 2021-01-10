package commands

import (
        "domain"
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

	//userid, err := repository.Mainrepo.StoreUser (cmd_newuser.P)
	_, err := repository.Mainrepo.StoreUser (cmd_newuser.P)
	if err != nil {panic (err)}
	logging.GetLogger("INFO").WithFields(logrus.Fields{
        "package": "commands",
        "source":"concreteCommands.go",
        "func": "CMD NewUser",
        }).Info("errtxt-command executed successfully")

	return nil

}

// New expense Command
type CMD_NewExpense struct {
	P *domain.Expense
}

func (cmd_newexpense CMD_NewExpense) Execute () error {
	rand.Seed(time.Now().UnixNano())
	cmd_newexpense.P.ExpenseId = rand.Intn (1000)

	_, _, err := repository.Mainrepo.StoreExpense (cmd_newexpense.P)
	if err != nil {panic (err)}
	logging.GetLogger("INFO").WithFields(logrus.Fields{
        "package": "commands",
        "source":"concreteCommands.go",
        "func": "CMD NewExpense",
        }).Info("errtxt-command executed successfully")
	
	return nil

}
