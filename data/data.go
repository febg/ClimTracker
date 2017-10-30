package data

import (
	"database/sql"
	"encoding/json"
	"log"
	"sync"

	"github.com/febg/Climbtracker/gym"
	"github.com/febg/Climbtracker/tools"
	"github.com/febg/Climbtracker/user"
	//"github.com/go-sql-driver/mysql" used as MySQL driver only
	_ "github.com/go-sql-driver/mysql"
)

// DConfig data-type is used for creating go routines groups
type DConfig struct {
	IG *sync.WaitGroup
}

// CheckUserExistance looks if client exists in users table
// func CheckUserExistance(DB *sql.DB, uData user.UserData) error {
// 	log.Printf("-> [INFO] Cheking user existance in data base...")
// 	err := validateRagistration(DB, uData)
// 	if err != nil {
// 		log.Printf("-> [ERROR] Unable to validate registration")
// 		return err
// 	}
// 	log.Printf("-> [INFO] User not registered, preparing to store data")
// 	return nil
// }

// RegisterUser handles the storage of new user in database
// func RegisterUser(DB *sql.DB, uData user.UserData) error {
// 	log.Printf("-> [INFO] Registering in data base...")
// 	err := sendUser(DB, uData)
// 	if err != nil {
// 		log.Printf("-> [ERROR] Unable to store user in database")
// 		return err
// 	}
// 	log.Printf("-> [INFO] User registered in database successfully")
// 	return nil
// }

// InitializeUserData creates a unique table for each user in database
func InitializeUserData(DB *sql.DB, uData user.UserData) {
	log.Printf("-> [INFO] Initializing user information")
	var initGroup sync.WaitGroup
	dataInit := DConfig{
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
func NewUser(DB *sql.DB, b []byte) (bool, error) {
	var uData user.UserData
	err := json.Unmarshal(b, &uData)
	if err != nil {
		log.Printf("-> [ERROR] Unable to Unmarshal user information: %v", err)
		return false, err
	}

	err = validateRagistration(DB, uData)
	if err != nil {
		log.Printf("-> [ERROR] Unable to validate registration")
		return true, err
	}

	err = sendUser(DB, uData)
	if err != nil {
		log.Printf("-> [ERROR] Unable to store user in database")
		return false, err
	}

	InitializeUserData(DB, uData)

	return true, nil
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
	dData := getClimbingData(DB, uID)
	cData := gym.ClimbingData{
		Data: dData,
	}
	return &cData, nil
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

func NewPullUp(DB *sql.DB, d []byte) error {
	var P user.NewPullUp
	err := json.Unmarshal(d, &P)
	if err != nil {
		log.Printf("-> [ERROR] Unable to Unmarshal pull up information: %v", err)
		return err
	}
	err = validateUID(DB, P.UserID)
	if err != nil {
		log.Printf("-> [ERROR] User not found: %v", err)
		return err
	}
	err = recordPullUp(DB, P)
	if err != nil {
		log.Printf("-> [ERROR] Unable to record block entry")
		return err
	}
	log.Printf("-> [INFO] Block recorded successfully")
	return nil
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

func GetFriends(DB *sql.DB, uID string) (*user.PrivateUser, error) {
	err := validateUID(DB, uID)
	if err != nil {
		log.Printf("-> [ERROR] User not found: %v", err)
		return nil, err
	}
	//TODO change: intead of privUser just slice.. make fIds private
	friendList := user.FriendList{}
	privUser := user.PrivateUser{}
	err = getFriendList(DB, uID, &friendList)
	if err != nil {
		log.Printf("-> [ERROR] Unable to get connections information")
		return nil, err
	}
	var friendsGroup sync.WaitGroup
	getFData := DConfig{
		IG: &friendsGroup,
	}
	for _, v := range friendList.Friends {
		friendsGroup.Add(1)
		go getFData.getPublicData(DB, v, &privUser)
	}
	friendsGroup.Wait()

	//err = getPublicData(DB, uID, privUser)
	return &privUser, nil
}

func (d *DConfig) getPublicData(DB *sql.DB, uID string, privUser *user.PrivateUser) error {
	defer d.IG.Done()
	pubUser := user.PublicUser{}
	err := getPublicProfile(DB, uID, &pubUser)
	if err != nil {
		log.Printf("-> [ERROR] Unable to get Public Profile: %v", err)
		return err
	}
	err = getClimbingStats(DB, uID, &pubUser)
	if err != nil {
		log.Printf("-> [ERROR] Unable to get Climbing Stats: %v", err)
		return err
	}
	privUser.FInfo = append(privUser.FInfo, pubUser)
	return nil
}
