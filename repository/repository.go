package repository

import (
	"fmt"
	"domain"
//	"errors"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)
var Mainrepo Repository

type Repository interface {
	StoreUser (user *domain.User) (userid int, err error)
	StoreExpense (expense *domain.Expense) (userId, expenseId int, err error)
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
	fmt.Println (sql_text)
	stmt, err := r.pool.Prepare(sql_text)
	if err != nil { panic(err) }
	defer stmt.Close()

	_, err = stmt.Exec (user.UserId, user.Name, user.Surname, user.Email, "20210109", user.Info, user.Gender)
	if err != nil { panic(err) }
	return user.UserId, nil
}
func (r *repoSqlite) StoreExpense (expense *domain.Expense) (userid, expenseId int, err error) {
	sql_text := "INSERT INTO EXPENSES (USERID,EXPENSEID,DATE,EXPITEM,AMOUNT,INFO) values(?, ?, ?, ?, ?, ?)"
	stmt, err := r.pool.Prepare(sql_text)
	if err != nil { panic(err) }
	defer stmt.Close()

	_, err = stmt.Exec (expense.UserId, expense.ExpenseId, "20210110", expense.ExpItem, expense.Amount,expense.Info)
	if err != nil { panic(err) }
	return 0,0,nil
}
func (r *repoSqlite) StoreSession(session *domain.Session) (sessionId int, err error) {
	// STORE SESSION TO USER TABLE
	return 0, nil
}
func (r *repoSqlite) GetUserByName(username string) (user *domain.User, err error) {
	// STORE SESSION TO USER TABLE
	return nil, nil
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
