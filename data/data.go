package data

import (
	"database/sql"
	"log"

	//"github.com/febg/Climbtracker/api"
	//"github.com/febg/Climbtracker/data"
	//"github.com/go-sql-driver/mysql" used as MySQL driver only
	_ "github.com/go-sql-driver/mysql"
)

// UserData represents the user climbing data from data base
type UserData struct {
	Name     string
	Email    string
	Password string
}

// CheckUserExistance looks if client exists in users table
func CheckUserExistance(DB *sql.DB, uData []byte) (bool, error) {
	log.Printf("-> [REQUEST] Cheking user existance in data base...")
	if success, err := ValidateRagistration(DB, uData); success != true {
		if err != nil {
			log.Printf("-> [ERROR] Unable to validate registration")
			return false, err
		}
		log.Printf("-> [ERROR] User already in data base")
		return false, nil
	}
	log.Printf("-> [REQUEST] User not registered, preparing to store data")
	return true, nil
}

// NewMySQL creates a connection to a MySQL database on AWS
func NewMySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", "awsuser:password@tcp(mydbinstance2.cr2vnklpvvxv.us-east-2.rds.amazonaws.com:3306)/mydb?charset=utf8")
	return db, err
}

// RegisterUser handles user registration in MySQL data base
func RegisterUser(DB *sql.DB, uData []byte) bool {

	return true
}

// NewLocalMySQL creates a connection to a MySQL database on local nerwork on port :3306
func NewLocalMySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:1692Ubc!@localhost:3306")
	return db, err
}
