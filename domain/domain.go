package domain

import (
	"time"
)
type User struct {
        UserId int `json:"userid"`
        Name string `json:"name"`
        Surname string `json:"surname"`
        Email string `json:"email"`
        DateCreated time.Time `json:"datecreated"`
        Info string `json:"info"`
        Gender string `json:"gender"`
}
type Expense struct {
        UserId int `json:"userid"`
        ExpenseId int `json: "expenseid"`
        Date time.Time `json: "date"`
        ExpItem string `json: "expitem"`
        Amount float64 `json: "amount"`
        Info string `json: "indo"`
}

type Session struct {
        SessionId int `json:"sessionid"`
        User *User `json:"user"`
        TimeOpened time.Time `json:"timeopened"`
        IdleTime time.Time `json:"idletime"`
        TimeClosed time.Time `json: "timeclosed"`
}
type Passwd struct {
        UserId int `json: "userid"`
        Passwd string `json: "passwd"`
}

type Userlog struct {
	userId int `json:"userid"`
	logDate time.Time `json:"logdate"`
	action string `json:"action"`
	params string `json:"params"`
}

