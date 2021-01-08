package commands

import (
//        "domain"
//        "errors"
//        "logging"
//        "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

// New user Command
type CMD_NewUserParams struct {
	UserId int `json:"userid"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	DateCreated time.Time `json:"datecreated"`
	Info string `json:"info"`
	Gender string `json:"gender"`
}

type CMD_NewUser struct {
	P *CMD_NewUserParams
}

func (cmd_newuser CMD_NewUser) Execute () error {
	cmd_newuser.P.UserId = rand.Intn (1000)
	cmd_newuser.P.DateCreated = time.Now ()

	return nil

	// Write it to dbase and put it active players list
}
