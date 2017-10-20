package data

import "database/sql"

// Data represents the user climbing data from data base
type Data interface {
}

//MySQLDB figure out later
type MySQLDB struct {
}

// CheckUserExistance looks if client exists in users table
func CheckUserExistance() {

}

// NewMySQL creates a connection to a MySQL database on AWS
func NewMySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", "awsuser:password@tcp(mydbinstance2.cr2vnklpvvxv.us-east-2.rds.amazonaws.com:3306)/mydb?charset=utf8")
	return db, err
}

// NewLocalMySQL creates a connection to a MySQL database on local nerwork on port :3306
func NewLocalMySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:1692Ubc!@localhost:3306")
	return db, err
}
