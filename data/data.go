package data

import "database/sql"

// Data represents the user climbing data from data base
type Data struct {
}

type MySQLDB struct {
}

// CheckUserExistance looks if client exists in users table
func (d *Data) CheckUserExistance() {

}

// NewMySQL creates a connection to a MySQL database on AWS
func (d *Data) NewMySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", "awsuser:password@tcp(mydbinstance2.cr2vnklpvvxv.us-east-2.rds.amazonaws.com:3306)/mydb?charset=utf8")
	return db, err
}
