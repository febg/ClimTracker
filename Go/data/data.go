package data

import (
	"database/sql"
	"encoding/json"
	"log"

	//"../tools"

	"github.com/febg/Climbtracker/Go/tools"
	//"github.com/go-sql-driver/mysql" used as MySQL driver only

	_ "github.com/go-sql-driver/mysql"
)

// UserData represents the user climbing data from data base
type UserData struct {
	Name     string
	Email    string
	Password string
	UserID   string
}

// CheckUserExistance looks if client exists in users table
func CheckUserExistance(DB *sql.DB, uData UserData) (bool, error) {
	log.Printf("-> [INFO] Cheking user existance in data base...")
	if success, err := validateRagistration(DB, uData); success != true {
		if err != nil {
			log.Printf("-> [ERROR] Unable to validate registration")
			return false, err
		}
		log.Printf("-> [ERROR] User already in data base")
		return false, nil
	}
	log.Printf("-> [INFO] User not registered, preparing to store data")
	return true, nil
}

// RegisterUser handles the storage of new user in database
func RegisterUser(DB *sql.DB, uData UserData) (bool, error) {
	log.Printf("-> [INFO] Registering in data base...")
	if success, err := sendUser(DB, uData); err != nil {
		log.Printf("-> [ERROR] Unable to store user in database")
		return success, err
	}
	log.Printf("-> [INFO] User registered in database successfully")
	return true, nil
}

// NewUserTable creates a unique table for each user in database
func NewUserTable(DB *sql.DB, uData UserData) (bool, error) {
	log.Printf("-> [INFO] Creating unique table for new user")
	if newTable, err := createUserTable(DB, uData); newTable != true {
		if err != nil {
			log.Printf("-> [ERROR] Unable to create table")
		}
	}
	log.Print("-> [INFO] Table created successfully")
	return true, nil
}

// NewUser handles user registration in MySQL data base
func NewUser(DB *sql.DB, uData []byte) (bool, error) {
	var data UserData
	err := json.Unmarshal(uData, &data)
	if err != nil {
		log.Printf("-> [ERROR] Unable to Unmarshal user information: %v", err)
		return false, err
	}
	if s1, err := CheckUserExistance(DB, data); err != nil {
		return false, err
	} else if s1 == true {
		if s2, err := RegisterUser(DB, data); s2 != true {
			if err != nil {
				return s2, err
			}
			return s2, nil
		}
		if s3, err := NewUserTable(DB, data); s3 != true {
			if err != nil {
				log.Printf("-> [ERROR] Unable to create new table")
				return false, err
			}
			return false, nil
		}
		return true, nil
	}
	return false, nil
}

// NewMySQL creates a connection to a MySQL database on AWS
func NewMySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", "awsuser:password@tcp(mydbinstance2.cr2vnklpvvxv.us-east-2.rds.amazonaws.com:3306)/mydb?charset=utf8")
	return db, err
}

// NewLocalMySQL creates a connection to a MySQL database on local nerwork on port :3306
func NewLocalMySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:1692Ubc!@tcp(localhost:3306)/test02?charset=utf8")
	return db, err
}

func InitializeTables(tables []string) (int, error) {

	return 0, nil
}

func LogIn(DB *sql.DB, uData []byte) (string, error) {
	var data UserData
	err := json.Unmarshal(uData, &data)
	if err != nil {
		log.Printf("-> [ERROR] Unable to Unmarshal user information: %v", err)
		return "", err
	}
	log.Printf("-> [LOG] Obtaining user's stored password")
	hpwd, uID, err := getUserPassword(DB, data)
	if err != nil {
		log.Printf("->[ERROR] Unable to obtained stored password: %v", err)
		return "", err
	}
	if !tools.ComparePasswords(data.Password, hpwd) {
		log.Printf("->[ERROR] Unable to Verify password")
		return "wpwd", nil
	}
	log.Printf("-> [INFO] User Authenticated successfully")
	return uID, nil
}

func ClimbingHistory(DB *sql.DB, uID string) (string, error) {
	getClimbingData(DB, uID)

	return "", nil

}
