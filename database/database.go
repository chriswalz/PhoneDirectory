package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type PhoneUser struct {
	First  string
	Last   string
	Ext    string
	Dept   string
	Number string
}

var db *sql.DB

func Prepare() {
	var err error
	db, err = sql.Open("mysql", "root:testpassword@tcp(127.0.0.1:3306)/hello")
	handlerError(err)
}

func GetRows() []*PhoneUser {
	var err error
	rows, err := db.Query("select first, last, ext, dept, number from phone")
	handlerError(err)
	defer rows.Close()
	users := make([]*PhoneUser, 0, 150)
	i := 0
	for rows.Next() {
		var user PhoneUser
		err = rows.Scan(&user.First, &user.Last, &user.Ext, &user.Dept, &user.Number)
		users = append(users, &user)
		handlerError(err)
		i++
	}
	err = rows.Err()
	handlerError(err)

	return users

}
func handlerError(err error) {
	if err != nil {
		panic(err)
	}
}
