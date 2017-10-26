package data

import (
	"database/sql"
	"encoding/json"
	"log"
	"sync"

	//"../tools"

	"github.com/febg/Climbtracker/gym"
	"github.com/febg/Climbtracker/tools"
	"github.com/febg/Climbtracker/user"

	//"github.com/go-sql-driver/mysql" used as MySQL driver only
	_ "github.com/go-sql-driver/mysql"
)

type DataConfig struct {
	IG *sync.WaitGroup
}

// CheckUserExistance looks if client exists in users table
func CheckUserExistance(DB *sql.DB, uData user.UserData) (bool, error) {
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
func RegisterUser(DB *sql.DB, uData user.UserData) (bool, error) {
	log.Printf("-> [INFO] Registering in data base...")
	if success, err := sendUser(DB, uData); err != nil {
		log.Printf("-> [ERROR] Unable to store user in database")
		return success, err
	}
	log.Printf("-> [INFO] User registered in database successfully")
	return true, nil
}

// InitializeUserData creates a unique table for each user in database
func InitializeUserData(DB *sql.DB, uData user.UserData) {
	log.Printf("-> [INFO] Initializing user information")
	var initGroup sync.WaitGroup
	dataInit := DataConfig{
		IG: &initGroup,
	}
	initGroup.Add(1)
	go dataInit.initializeUserTable(DB, uData)
	initGroup.Add(1)
	go dataInit.initializeClimbingstats(DB, uData)
	initGroup.Add(1)
	go dataInit.initializePullUp(DB, uData)
	initGroup.Wait()

	log.Print("-> [INFO] User information initialized")
	return
}

// NewUser handles user registration in MySQL data base
func NewUser(DB *sql.DB, uData []byte) (bool, error) {
	var data user.UserData
	err := json.Unmarshal(uData, &data)
	if err != nil {
		log.Printf("-> [ERROR] Unable to Unmarshal user information: %v", err)
		return false, err
	}
	if s1, err := CheckUserExistance(DB, data); err != nil {
		return false, err
	} else if s1 == true {
		data.QrCode = tools.GenerateQrCode(data.Email)
		if s2, err := RegisterUser(DB, data); s2 != true {
			if err != nil {
				return s2, err
			}
			return s2, nil
		}
		InitializeUserData(DB, data)
		return true, err
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
	db, err := sql.Open("mysql", "root:1692Ubc!@tcp(localhost:3306)/test3?charset=utf8")
	return db, err
}

func InitializeTables(tables []string) (int, error) {

	return 0, nil
}

func LogIn(DB *sql.DB, uData []byte) (string, error) {
	var data user.UserData
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

// ClimbingHistory gets all climbing history for user
func ClimbingHistory(DB *sql.DB, uID string) (*gym.ClimbingData, error) {
	cData := getClimbingData(DB, uID)

	return cData, nil
}

func CheckIn(DB *sql.DB, d []byte) error {
	var C user.NewCheckIn
	err := json.Unmarshal(d, &C)
	if err != nil {
		log.Printf("-> [ERROR] Unable to Unmarshal user information: %v", err)
		return err
	}
	err = validateUID(DB, C.UserID)
	if err != nil {
		log.Printf("-> [ERROR] User not found: %v", err)
		return err
	}
	err = recordBlock(DB, C)
	if err != nil {
		log.Printf("-> [ERROR] Unable to record block entry")
		return err
	}
	log.Printf("-> [INFO] Block recorded successfully")
	return nil
}

func FriendRequest(DB *sql.DB, uID string, femail string) error {
	err := validateUID(DB, uID)
	if err != nil {
		log.Printf("-> [ERROR] User not found: %v", err)
		return err
	}

	fuID, err := validateFriendInfo(DB, femail)
	if err != nil {
		log.Printf("-> [ERROR] User not found: %v", err)
		return err
	}

	err = checkUsersConnection(DB, uID, fuID)
	if err != nil {
		log.Printf("-> [ERROR] Friendship connection check: %v", err)
		return err
	}
	err = createUsersConnection(DB, uID, fuID)
	if err != nil {
		log.Printf("-> [ERROR] Friendship connection: %v", err)
	}
	log.Printf("-> [INFO] Friendship connection recorded successfully")
	return nil
}

func GetFriends(DB *sql.DB, uID string) error {
	err := validateUID(DB, uID)
	if err != nil {
		log.Printf("-> [ERROR] User not found: %v", err)
		return err
	}

	return nil
}
