package data

import (
	"database/sql"
	"errors"
	"log"

	"github.com/febg/Climbtracker/gym"
	"github.com/febg/Climbtracker/tools"
)

//ValidateRagistration executes MySQL query to check user email in database
func validateRagistration(DB *sql.DB, uData UserData) (bool, error) {
	rows, err := DB.Query(`SELECT Email FROM UserInformation WHERE Email=` + tools.QueryField(uData.Email) + `;`)
	if err != nil {
		log.Printf("-> [ERROR] Validation Query: %v, %v", err, DB)
		return false, err
	}
	defer rows.Close()
	var email string
	for rows.Next() {
		err := rows.Scan(&email)
		if err != nil {
			log.Printf("-> [ERROR] SQL response: %v", err)
			return false, err
		}
	}
	if email == "" {
		return true, nil
	}

	return false, nil
}
func validateUID(DB *sql.DB, uID string) error {
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

	return nil
}

// SendUser prepares and executes MySQL query to store user in data base
func sendUser(DB *sql.DB, uData UserData) (bool, error) {
	log.Printf("PWD L: %v", len(uData.Password))

	myQuery := `INSERT INTO UserInformation VALUES (NULL,` + `'` + tools.GetDate() + `','` + uData.Name + `','` + uData.Password + `','` + uData.Email + `','` + uData.UserID + `');`
	stmt, err := DB.Prepare(myQuery)
	if err != nil {
		log.Printf("->[ERROR] Registration query preparation: %v", err)
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Printf("->[ERROR] Registration query execution: %v", err)
		return false, err
	}
	return true, nil
}

func getUserPassword(DB *sql.DB, uData UserData) (string, string, error) {
	rows, err := DB.Query(`SELECT password, uID FROM UserInformation WHERE Email="` + uData.Email + `";`)
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

func getClimbingData(DB *sql.DB, uID string) *gym.ClimbingData {
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
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
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

		cData.Append(dData)

	}

	return &cData
}

func recordBlock(DB *sql.DB, cData NewCheckIn) error {
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

func (d *DataConfig) initializeUserTable(DB *sql.DB, uData UserData) {
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

func (d *DataConfig) initializeClimbingstats(DB *sql.DB, uData UserData) {
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

func (d *DataConfig) initializePullUp(DB *sql.DB, uData UserData) {
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
