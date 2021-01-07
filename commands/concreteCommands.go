package commands

import (
        "domain"
        "errors"
        "logging"
        "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

// New user Command
type CMD_NewUser struct {
	Name, Surname, Email, Info, Gender string
}

func (cmd_newuser *CMD_NewUser) Execute () error {
	user := new (domain.user)
	user.name = Name
	user.surname = Surname
	user.email = Email
	user.info = Info
	user.gender = Gender
	user.userId = rand.intn (1000)
	user.dateCreated = time.Now ()

	return nil

	// Write it to dbase and put it active players list
}

