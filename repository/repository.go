package repository

import (
//	"fmt"
	"strconv"
	"domain"
	"errors"
	"logging"
	"time"
	"github.com/sirupsen/logrus"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)
var Mainrepo Repository

type Repository interface {
	StoreUser (user *domain.User) (userid int, err error)
	StoreExpense (expense *domain.Expense) (expenseId int, err error)
	StoreSession (session *domain.Session) (sesssionId int, err error)
	GetUserByName (username string) (user *domain.User, err error)
	GetUserById (userId int) (user *domain.User, err error)
	GetExpenseById (expenseId int) (expense *domain.Expense, err error) 
	GetExpensesOfUser (user *domain.User, startdate, enddate string) (expenses []*domain.Expense, err error) 
}

type repoSqlite struct {
	pool *sql.DB
}

//type repoInmem struct {
//	pool *inmemhendle
//}

func NewSqliteRepository (p *sql.DB) Repository {
	return &repoSqlite {
		pool : p,
	}
}

//func NewInmemRepository (p *inmemhandle) Repository {
//	return &repoInmem {
//		pool : p,
//	}
//}

//SQLITE 
func (r *repoSqlite) StoreUser(user *domain.User) (userid int, err error) {
	sql_text := "INSERT INTO USER (USERID,NAME,SURNAME,EMAIL,DATECREATED,INFO,GENDER) values(?, ?, ?, ?, ?, ?, ?)"
	stmt, err := r.pool.Prepare(sql_text)
	if err != nil {
		logging.GetLogger("DEFAULT").WithFields(logrus.Fields{
		"package":"repository",
		"source":"repository.go",
		"func":"StoreUser",
		"errorpoint":"sqlPrepare",
		"sqltext":sql_text,
		}).Error (err.Error())
		return user.UserId, errors.New ("Kullanici sisteme kayit edilemedi ")
	}
	defer stmt.Close()

	t := time.Now()
	_, err = stmt.Exec (user.UserId, user.Name, user.Surname, user.Email, t.Format("20060102"), user.Info, user.Gender)
	if err != nil {
		logging.GetLogger("DEFAULT").WithFields(logrus.Fields{
		"package":"repository",
		"source":"repository.go",
		"func":"StoreUser",
		"errorpoint":"sqlExec",
		"sqltext":sql_text,
		}).Error (err.Error())
		return user.UserId, errors.New ("Kullanici sisteme kayit edilemedi ")
	}

	logging.GetLogger("INFO").WithFields(logrus.Fields{
	"package":"repository",
	"source":"repository.go",
	"func":"StoreUser",
	"userId": user.UserId,
	}).Info ("Kullanici basari ile kaydedildi. Kullanici ID :", strconv.Itoa(user.UserId))
	return user.UserId, nil
}
func (r *repoSqlite) StoreExpense (expense *domain.Expense) (expenseId int, err error) {
	sql_text := "INSERT INTO EXPENSES (USERID,EXPENSEID,DATE,EXPITEM,AMOUNT,INFO) values(?, ?, ?, ?, ?, ?)"
	stmt, err := r.pool.Prepare(sql_text)
	if err != nil {
		logging.GetLogger("DEFAULT").WithFields(logrus.Fields{
		"package":"repository",
		"source":"repository.go",
		"func":"StoreExpense",
		"errorpoint":"sqlPrepare",
		"sqltext":sql_text,
		}).Error (err.Error())
		return expense.ExpenseId, errors.New ("Harcama sisteme kayit edilemedi ")
	}
	defer stmt.Close()

	t := time.Now()
	_, err = stmt.Exec (expense.UserId, expense.ExpenseId, t.Format("20060102"), expense.ExpItem, expense.Amount,expense.Info)
	if err != nil {
		logging.GetLogger("DEFAULT").WithFields(logrus.Fields{
		"package":"repository",
		"source":"repository.go",
		"func":"StoreExpense",
		"errorpoint":"sqlExec",
		"sqltext":sql_text,
		}).Error (err.Error())
		return expense.ExpenseId, errors.New ("Harcama sisteme kayit edilemedi ")
	}
	logging.GetLogger("INFO").WithFields(logrus.Fields{
	"package":"repository",
	"source":"repository.go",
	"func":"StoreExpense",
	"userId": expense.ExpenseId,
	}).Info ("Expense Basari ile kaydedildi. expense ID :", expense.ExpenseId)
	return expense.ExpenseId ,nil
}
func (r *repoSqlite) StoreSession(session *domain.Session) (sessionId int, err error) {
	// STORE SESSION TO USER TABLE
	return 0, nil
}
func (r *repoSqlite) GetUserByName(username string) (user *domain.User, err error) {
	rows, err := r.pool.Query("SELECT USERID,NAME,SURNAME,EMAIL,INFO,GENDER FROM USER ORDER BY NAME WHERE NAME = ?", user.Name)
	if err != nil {
		logging.GetLogger("DEFAULT").WithFields(logrus.Fields{
		"package":"repository",
		"source":"repository.go",
		"func":"GetUserByName",
		"errorpoint":"sqlQuery",
		"username":username,
		}).Error (err.Error())
		return nil, errors.New ("Kullanici tabloda aranirken hata olustu")
	}
	defer rows.Close()

	var tmpuser domain.User
	rowcount := 0
	for rows.Next () {
		if rowcount > 0 { // more than one user with the same name
			logging.GetLogger("DEFAULT").WithFields(logrus.Fields{
			"package":"repository",
			"source":"repository.go",
			"func":"GetUserByName",
			"errorpoint":"",
			"username":username,
			}).Error ("Ayni Isimde birden fazla kullanici var")
			return nil, errors.New ("Ayni isimde birden fazla kullnici var. Kullanic isimleri unique olmalidir")
		}//rowcount==0

		if err := rows.Scan (&tmpuser.UserId,
				&tmpuser.Name,
				&tmpuser.Surname,
				&tmpuser.Email,
				&tmpuser.Info,
				&tmpuser.Gender); err != nil {
			logging.GetLogger("DEFAULT").WithFields(logrus.Fields{
			"package":"repository",
			"source":"repository.go",
			"func":"GetUserByName",
			"errorpoint":"rows.Scan",
			"username":username,
			}).Error (err.Error())
			return nil, errors.New ("Rows scan edilirken hata olustu:" + username)
		} else { // we got the user
			rowcount += 1
			logging.GetLogger("INFO").WithFields(logrus.Fields{
			"package":"repository",
			"source":"repository.go",
			"func":"GetUserByName",
			"userId": tmpuser.UserId,
			"username":tmpuser.Name,
			"usersurname":tmpuser.Surname,
			}).Info (tmpuser.Name, " isimli kullanici bulundu")
		}
	}

	return &tmpuser, nil
}
func (r *repoSqlite) GetUserById(userid int) (user *domain.User, err error) {
	// STORE SESSION TO USER TABLE
	return nil, nil
}
func (r *repoSqlite) GetExpenseById (expenseId int) (expense *domain.Expense, err error) {
	// GET EXPENSE
	return nil, nil
}
func (r *repoSqlite) GetExpensesOfUser (user *domain.User, startdate, enddate string) (expenses []*domain.Expense, err error) {
	// GET EXPENSES
	exp := make ([]*domain.Expense, 10)
	return exp, nil
}
/*
//INMEM
func (r *repoInmem) StoreUser(user *domain.User) (userid int64, err error) {
	// STORE USER TO USER TABLE
	return 0, nil
}

func (r *repoInmem) StoreExpense (user *domain, expense *domain) (userid, expenseId int64, err error) {
	//STORE EXPENSE TO EXPENSE TABLE
	return 0,0,nil
}

func (r *repoInmem) StoreSession(session *domain.Session) (sessionId int64, err error) {
	// STORE SESSION TO USER TABLE
	return 0, nil
}
func (r *repoInmem) GetUserByName(username string) (user *domain.User, err error) {
	// GET USER
	return nil, nil
}
func (r *repoInmem) GetUserById(userid int64) (user *domain.User, err error) {
	// GET USER
	return nil, nil
}
func (r *repoInmem) GetExpensesOfUser (user *domain.User, startdate, enddate string) (expenses []*domain.Expense, err error) {
	// GET EXPENSES
	return []nil, nil
}
*/
