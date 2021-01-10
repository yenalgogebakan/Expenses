module github.com/yenalgogebakan/Expenses

replace logging => ./logging

replace commands => ./commands

replace domain => ./domain

replace repository => ./repository

go 1.15

require (
	commands v0.0.0-00010101000000-000000000000 // indirect
	domain v0.0.0-00010101000000-000000000000 // indirect
	github.com/gin-gonic/gin v1.6.3 // indirect
	github.com/mattn/go-sqlite3 v1.14.6 // indirect
	github.com/sirupsen/logrus v1.7.0 // indirect
	logging v0.0.0-00010101000000-000000000000 // indirect
	repository v0.0.0-00010101000000-000000000000 // indirect
)
