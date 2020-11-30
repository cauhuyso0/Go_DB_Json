package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//Phương thức kết nối với db mysql
func DBConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "posts"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return
}
