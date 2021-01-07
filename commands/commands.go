package commands

import (
	"errors"
	"logging"
	"github.com/sirupsen/logrus"
)

type Command interface {
	Execute () error
}

type Invoker struct {
	commands map[string]*Command
}

func (invoker *Invoker) RegCommand (commandName string, c *Command) error {
	if invoker.commands == nil {
		invoker.commands = make (map[string]*Command)
		logging.GetLogger("INFO").WithFields(logrus.Fields{
                "package": "commands",
		"source":"comands.go",
                "func": "Invoker regCommand",
	        }).Info("errtxt-commands map is initialized with make")
	}
	if _, ok := invoker.commands[commandName]; ok {
		logging.GetLogger("DEFAULT").WithFields(logrus.Fields{
                "package": "commands",
		"source": "commands.go",
                "func": "Invoker regCommand",
		"commandname" : commandName,
	        }).Error("errtxt-Attemp to register a preregistered command")
		return errors.New ("RegCommand : Attemp to register a preregistered command")
	}
	invoker.commands[commandName] = c
	logging.GetLogger("INFO").WithFields(logrus.Fields{
        "package": "commands",
	"source": "commands.go",
        "func": "Invoker regCommand",
	"commandname" : commandName,
	}).Info("errtxt-command is registered into command map")

	return nil
}

func (invoker *Invoker) ExeCommands (commandName string) error {
// If commandname == "", execute all commands, else the specic command
	if commandName == "" {
		for cmdname, cmd := range invoker.commands {
			if err := *cmd.Execute (); err != nil {
				logging.GetLogger("DEFAULT").WithFields(logrus.Fields{
				"package": "commands",
				"source": "commands.go",
				"func": "Invoker exeCommands",
				"commandname" : cmdname,
				"goterr": err.Error(),
				}).Error("errtxt-exec command returned error")
				return errors.New ("exeCommands : error in executing command")
			} else {
				logging.GetLogger("INFO").WithFields(logrus.Fields{
				"package": "commands",
				"source": "commands.go",
				"func": "Invoker exeCommands",
				"commandname" : cmdname,
				}).Info("Command executed successfully")
			}
		}

		return nil
	} else { //specific command
		if cmd, ok := invoker.commands[commandName]; ok { // command exists
			if err := cmd.Execute (); err != nil {
				logging.GetLogger("DEFAULT").WithFields(logrus.Fields{
				"package": "commands",
				"source": "commands.go",
				"func": "Invoker exeCommands",
				"commandname" : commandName,
				"goterr": err.Error(),
				}).Error("errtxt-exec command returned error")
				return errors.New ("exeCommands :exec command returned error")
			} else {
				logging.GetLogger("INFO").WithFields(logrus.Fields{
				"package": "commands",
				"source": "commands.go",
				"func": "Invoker exeCommands",
				"commandname" : commandName,
				}).Info("Command executed successfully")
				return nil
			}
		} else { //No such command is registered
			logging.GetLogger("DEFAULT").WithFields(logrus.Fields{
			"package": "commands",
			"source": "commands.go",
			"func": "Invoker exeCommands",
			"commandname" : commandName,
			}).Error("errtxt-can not execute a non registered command")
			return errors.New ("errtxt-can not execute a non registered command")

		}
		return nil
	}
}





