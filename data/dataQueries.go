package data

import (
	"database/sql"
	"errors"
	"log"

	"github.com/febg/Climbtracker/gym"
	"github.com/febg/Climbtracker/tools"
	"github.com/febg/Climbtracker/user"
)

//ValidateRagistration executes MySQL query to check user email in database
func validateRagistration(DB *sql.DB, uData user.UserData) error {
	log.Printf("-> [INFO] Cheking user existance in data base...")
	rows, err := DB.Query(`SELECT Email FROM UserInformation WHERE Email=` + tools.QueryField(uData.Email) + `;`)
	if err != nil {
		log.Printf("-> [ERROR] Validation Query: %v", err)
		return err
	}
	defer rows.Close()
	var email string
	for rows.Next() {
		err := rows.Scan(&email)
		if err != nil {
			log.Printf("-> [ERROR] SQL response: %v", err)
			return err
		}
	}
	if email == "" {
		log.Printf("-> [INFO] User not registered, preparing to store data")
		return nil
	}
	log.Printf("-> [ERROR] User already registered in data base")
	return errors.New("User already registered in data base")
}

// SendUser prepares and executes MySQL query to store user in data base
func sendUser(DB *sql.DB, uData user.UserData) error {
	log.Printf("-> [INFO] Registering in data base...")
	//myQuery := `INSERT INTO UserInformation VALUES (NULL,` + `'` + tools.GetDate() + `','` + uData.Name + `','` + uData.Password + `','` + uData.Email + `','` + uData.UserID + `', 1);`
	stmt, err := DB.Prepare(`INSERT INTO UserInformation VALUES (NULL,` + `'` + tools.GetDate() + `','` + uData.Name + `','` + uData.Password + `','` + uData.Email + `','` + uData.UserID + `', 1);`)
	if err != nil {
		log.Printf("->[ERROR] Registration query preparation: %v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Printf("->[ERROR] Registration query execution: %v", err)
		return err
	}
	log.Printf("-> [INFO] User registered in database successfully")
	return nil
}

func getUserPassword(DB *sql.DB, uData user.UserData) (string, string, error) {
	rows, err := DB.Query(`SELECT password, uID FROM UserInformation WHERE Email=` + tools.QueryField(uData.Email) + `;`)
	if err != nil {
		log.Printf("-> [ERROR] Get user pwd query: %v", err)
		return "", "", err
	}
	defer rows.Close()
	var pass string
	var id string

	for rows.Next() {
		err = rows.Scan(&pass, &id)
		if err != nil {
			log.Printf("-> [ERROR] SQL response: %v", err)
			return "", "", err
		}

	}
	if pass == "" || id == "" {
		return "", "", nil
	}
	return pass, id, err
}

func getPublicProfile(DB *sql.DB, uID string, pubUser *user.PublicUser) error {
	log.Printf("-> [INFO] Obtaining %v Public Profile..", uID)
	rows, err := DB.Query(`SELECT Name, Email, Public FROM UserInformation WHERE uID =` + tools.QueryField(uID) + `;`)
	if err != nil {
		log.Printf("-> [ERROR] public profile query: %v", err)
		return err
	}
	defer rows.Close()
	var name string
	var email string
	var public string

	for rows.Next() {
		err = rows.Scan(&name, &email, &public)
		if err != nil {
			log.Printf("-> [ERROR] Publi Profile SQL response: %v", err)
			return err
		}

	}
	if name == "" || email == "" || public == "" {
		return errors.New("Public Profile Data not complete")
	}

	pubUser.Name = name
	pubUser.Email = email
	pubUser.Public = public
	log.Printf("-> [INFO] Public Profile successfully obtained")

	return nil
}

func getFriendList(DB *sql.DB, uID string, privUser *user.FriendList) error {
	log.Printf("-> [INFO] Obtaining friend list..")
	rows, err := DB.Query(`SELECT uID2 FROM UsersConnections WHERE uID1 = ` + tools.QueryField(uID) + `;`)
	if err != nil {
		log.Printf("-> [ERROR] Get friends data query: %v", err)
		return err
	}
	defer rows.Close()
	var uID2 string
	for rows.Next() {
		err = rows.Scan(&uID2)
		if err != nil {
			log.Printf("-> [ERROR] SQL response: %v", err)
			return err
		}
		privUser.Friends = append(privUser.Friends, uID2)
	}
	log.Printf("-> [INFO] Friend list obtained successfully..")
	return nil
}

func getClimbingData(DB *sql.DB, uID string) []gym.DayData {
	rows, err := DB.Query(`SELECT * FROM ClimbingSessions WHERE uID = ` + tools.QueryField(uID) + `;`)
	if err != nil {
		log.Printf("-> [ERROR] Get user climbing data query: %v", err)
		return nil
	}
	defer rows.Close()
	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		log.Printf("-> [ERROR] Get user climbing data query: %v", err)
		return nil // proper error handling instead of panic in your app
	}
	values := make([]sql.RawBytes, len(columns))
	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	cData := gym.ClimbingData{}
	data := []gym.DayData{}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil // proper error handling instead of panic in your app
		}

		dData := gym.DayData{
			Index: string(values[0]),
			Date:  string(values[1]),
			UId:   string(values[2]),
			V1:    string(values[3]),
			V2:    string(values[4]),
			V3:    string(values[5]),
			V4:    string(values[6]),
			V5:    string(values[7]),
			V6:    string(values[8]),
		}

		data = append(data, dData)
		cData.Append(dData)

	}
	return data
}

func getClimbingStats(DB *sql.DB, uID string, pubUser *user.PublicUser) error {
	log.Printf("-> [INFO] Obtaining %v Climbing Stats..", uID)
	rows, err := DB.Query(`SELECT * FROM ClimbingStats WHERE uID = ` + tools.QueryField(uID) + `;`)
	if err != nil {
		log.Printf("-> [ERROR] Get user climbing data query: %v", err)
		return nil
	}
	defer rows.Close()
	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		log.Printf("-> [ERROR] Get user climbing data query: %v", err)
		return nil // proper error handling instead of panic in your app
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	dData := gym.OverallData{}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil // proper error handling instead of panic in your app
		}
		dData = gym.OverallData{
			SDate: string(values[1]),
			MV1:   string(values[3]),
			MV2:   string(values[4]),
			MV3:   string(values[5]),
			MV4:   string(values[6]),
			MV5:   string(values[7]),
			MV6:   string(values[8]),
			Total: string(values[9]),
		}
	}
	err = getPullUpStats(DB, uID, &dData)
	if err != nil {
		log.Printf("-> [ERROR] Unable to obtain Pullup stats for user: %v", uID)
		pubUser.Climbing = dData
		return err
	}
	pubUser.Climbing = dData
	if dData.SDate == "" {
		log.Printf("-> [ERROR] Unable to obtain climbing stats for user: %v", uID)
		return errors.New("Climbing Stats Records Do Not Exist")
	}
	log.Printf("-> [INFO] Climbings stats Obtained Successfully")
	return nil
}

func getPullUpStats(DB *sql.DB, uID string, oD *gym.OverallData) error {
	log.Printf("-> [INFO] Obtaining %v PullUp Stats..", uID)
	rows, err := DB.Query(`SELECT Date, Count, Max FROM PullUpDB WHERE uID =` + tools.QueryField(uID) + `;`)
	if err != nil {
		log.Printf("-> [ERROR] Get Pullup Stats query: %v", err)
		return err
	}
	defer rows.Close()
	var date string
	var count string
	var max string

	for rows.Next() {
		err = rows.Scan(&date, &count, &max)
		if err != nil {
			log.Printf("-> [ERROR] Get Pullup Stats SQL response: %v", err)
			return err
		}

	}
	if date == "" || count == "" || max == "" {
		return errors.New("Pull-Up Data not complete")
	}

	oD.PDate = date
	oD.PCount = count
	oD.PMax = max
	return nil
}

func validateUID(DB *sql.DB, uID string) error {
	log.Printf("-> [INFO] Verifying user ID")
	rows, err := DB.Query(`SELECT Email FROM UserInformation WHERE uID=` + tools.QueryField(uID) + `;`)
	if err != nil {
		log.Printf("-> [ERROR] uID validation Query: %v, %v", err, DB)
		return err
	}
	defer rows.Close()
	var id string
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Printf("-> [ERROR] uID SQL response: %v", err)
			return err
		}
	}
	if id == "" {
		return errors.New("Usernot found in database")
	}
	log.Printf("-> [INFO] User ID verified")
	return nil
}

func recordBlock(DB *sql.DB, cData user.NewCheckIn) error {
	stmt, err := DB.Prepare(`INSERT INTO ClimbingSessions (` + tools.QueryTable("index") + `, ` + tools.QueryTable("Date") + `, ` + tools.QueryTable("uID") + `, ` + tools.QueryTable(tools.Boulder(cData.Level)) + `) VALUES (NULL,` + tools.QueryField(tools.GetDate()) + `,` + tools.QueryField(cData.UserID) + `, 1);`)
	if err != nil {
		log.Printf("-> [ERROR] Record Block query preparation: %v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Printf("-> [ERROR] Record Block query execution: %v", err)
		return err
	}
	log.Printf("-> [INFO] Block entry recorded successfully")
	return nil
}

func updateStats(DB *sql.DB, u user.NewCheckIn) {
	return
}

//TODO Insert Update
func recordPullUp(DB *sql.DB, p user.NewPullUp) error {
	// stmt, err := DB.Prepare(`INSERT INTO PullUpDB (` + tools.QueryTable("index") + `, ` + tools.QueryTable("Date") + `, ` + tools.QueryTable("uID") + `, ` + tools.QueryTable(tools.Boulder(cData.Level)) + `) VALUES (NULL,` + tools.QueryField(tools.GetDate()) + `,` + tools.QueryField(cData.UserID) + `, 1);`)
	// if err != nil {
	// 	log.Printf("-> [ERROR] Record Block query preparation: %v", err)
	// 	return err
	// }
	// defer stmt.Close()
	// _, err = stmt.Exec()
	// if err != nil {
	// 	log.Printf("-> [ERROR] Record Block query execution: %v", err)
	// 	return err
	// }
	// log.Printf("-> [INFO] Block entry recorded successfully")
	return nil
}

func createUsersConnection(DB *sql.DB, uID string, fInfo string) error {
	myQuery := `INSERT INTO UsersConnections VALUES (NULL,` + tools.QueryField(tools.GetDate()) + `,` + tools.QueryField(uID) + `,` + tools.QueryField(fInfo) + `);`
	stmt, err := DB.Prepare(myQuery)
	if err != nil {
		log.Printf("->[ERROR] Creating users connection query preparation: %v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Printf("->[ERROR] Creating users connection query execution: %v", err)
		return err
	}
	myQuery = `INSERT INTO UsersConnections VALUES (NULL,` + tools.QueryField(tools.GetDate()) + `,` + tools.QueryField(fInfo) + `,` + tools.QueryField(uID) + `);`
	stmt, err = DB.Prepare(myQuery)
	if err != nil {
		log.Printf("->[ERROR] Creating users connection query preparation: %v", err)
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Printf("->[ERROR] Creating users connection query execution: %v", err)
		return err
	}
	return nil
}

func checkUsersConnection(DB *sql.DB, uID1 string, uID2 string) error {
	query := `SELECT Date FROM UsersConnections WHERE uID1 = ` + tools.QueryField(uID1) + ` AND uID2 = ` + tools.QueryField(uID2) + `;`
	rows, err := DB.Query(query)
	if err != nil {
		log.Printf("-> [ERROR] uID validation Query: %v, %v", err, DB)
		return err
	}
	defer rows.Close()
	var date string
	for rows.Next() {
		err = rows.Scan(&date)
		if err != nil {
			log.Printf("-> [ERROR] Friendship connection SQL response: %v", err)
			return err
		}
	}
	if date == "" {
		return nil
	}
	return errors.New("Friend Ship Exists Created on:" + date)
}

func validateFriendInfo(DB *sql.DB, email string) (string, error) {
	log.Printf("-> [INFO] Verifiying requested user ID")
	rows, err := DB.Query(`SELECT uID FROM UserInformation WHERE Email=` + tools.QueryField(email) + `;`)
	if err != nil {
		log.Printf("-> [ERROR] Validation Query: %v, %v", err, DB)
		return "", err
	}
	defer rows.Close()
	var uID string
	for rows.Next() {
		err := rows.Scan(&uID)
		if err != nil {
			log.Printf("-> [ERROR] SQL response: %v", err)
			return "", err
		}
	}
	if uID == "" {
		return "", errors.New("validateFriendInfo: User does not exists")
	}
	log.Printf("-> [INFO] User ID obtained and verified")
	return uID, nil
}

func (d *DConfig) initializeUserTable(DB *sql.DB, uData user.UserData) {
	defer d.IG.Done()
	myquery := `INSERT INTO ClimbingSessions (` + tools.QueryTable("index") + `, ` + tools.QueryTable("Date") + `, ` + tools.QueryTable("uID") + `) VALUES (NULL,` + tools.QueryField(tools.GetDate()) + `,` + tools.QueryField(uData.UserID) + `);`
	stmt, err := DB.Prepare(myquery)
	if err != nil {
		log.Printf("-> [ERROR] Initialize Table query preparation: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Printf("-> [ERROR] Initialize Table query execution: %v", err)
		return
	}

	return
}

func (d *DConfig) initializeClimbingstats(DB *sql.DB, uData user.UserData) {
	defer d.IG.Done()
	myquery := `INSERT INTO ClimbingStats (` + tools.QueryTable("index") + `, ` + tools.QueryTable("Date") + `, ` + tools.QueryTable("uID") + `) VALUES (NULL,` + tools.QueryField(tools.GetDate()) + `,` + tools.QueryField(uData.UserID) + `);`
	stmt, err := DB.Prepare(myquery)
	if err != nil {
		log.Printf("-> [ERROR] Climbing stats query preparation: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Printf("-> [ERROR] Climbing stats query execution: %v", err)
		return
	}
	return
}

func (d *DConfig) initializePullUp(DB *sql.DB, uData user.UserData) {
	defer d.IG.Done()
	myquery := `INSERT INTO PullUpDB (` + tools.QueryTable("index") + `, ` + tools.QueryTable("Date") + `, ` + tools.QueryTable("uID") + `) VALUES (NULL, ` + tools.QueryField(tools.GetDate()) + `, ` + tools.QueryField(uData.UserID) + `);`
	stmt, err := DB.Prepare(myquery)
	if err != nil {
		log.Printf("-> [ERROR] PullUpDB query preparation: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Printf("-> [ERROR] PullUpDB query execution: %v", err)
		return
	}
	return
}
