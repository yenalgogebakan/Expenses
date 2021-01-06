package domain

import (
	"time"
)


type user struct {
	userId int64 		`json: "userid"`
	name string 		`json: "name"`
	surname string 		`json: "surname"`
	email string 		`json: "userid"`
	dateCreated time.Time 	`json: "datecreated"`
	info string 		`json: "info"`
	gender string 		`json: "gender"`
}

type passwd struct {
	userId int64		`json: "userid"`
	passwd string		`json: "passwd"`
}

type userlog struct {
	userId int64		`json: "userid"`
	logDate time.Time	`json: "logdate"`
	action string		`json: "action"`
	params string
}

type expenses struct {
	userId int64		`json: "userid"`
	enxpenseId int64	`json: "expenseid"`
	date time.Time		`json: "date"`
	expItem string		`json: "expitem"`
	amaunt float64		`json: "amaunt"`
	info string		`json: "indo"`
}

type sessions struct {
	sessionId int64		`json: "sessionid"`
	user *user		`json: "user"`
	timeOpened time.Time	`json: "timeopened"`
	idleTime time.Time	`json: "idletime"`
	timeClosed time.Time	`json: "timeclosed"`
}

