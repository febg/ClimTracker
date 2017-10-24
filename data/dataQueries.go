package data

import (
	"database/sql"
	"log"

	"github.com/febg/Climbtracker/gym"
	"github.com/febg/Climbtracker/tools"
)

//ValidateRagistration executes MySQL query to check user email in database
func validateRagistration(DB *sql.DB, uData UserData) (bool, error) {
	qe := "'" + uData.Email + "'"
	rows, err := DB.Query(`SELECT email FROM UserInformation WHERE Email=` + qe + `;`)
	if err != nil {
		log.Printf("-> [ERROR] Validation Query: %v, %v", err, DB)
		return false, err
	}
	defer rows.Close()
	var name string
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Printf("-> [ERROR] SQL response: %v", err)
			return false, err
		}
	}
	if name == "" {
		return true, nil
	}

	return false, nil
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
	rows, err := DB.Query(`SELECT password, tableID FROM users WHERE email="` + uData.Email + `";`)
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
			UId:  string(values[0]),
			Date: string(values[1]),
			V1:   string(values[2]),
			V2:   string(values[3]),
			V3:   string(values[4]),
			V4:   string(values[5]),
			V5:   string(values[6]),
			V6:   string(values[7]),
		}

		cData.Append(dData)

	}

	return &cData
}

func recordBlock(DB *sql.DB, cData NewCheckIn) error {
	stmt, err := DB.Prepare(`UPDATE ` + tools.QueryTable(cData.UserID) + ` SET ` + tools.Boulder(cData.Level) + `=` + tools.Boulder(cData.Level) + `+ 1 WHERE date=` + tools.QueryField(tools.GetDate()) + `;`)
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
	myquery := `INSERT INTO ClimbingSessions VALUES (NULL,` + tools.QueryField(tools.GetDate()) + `,` + tools.QueryField(uData.UserID) + `, 0,  0,  0,  0, 0, 0);`
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
	myquery := `INSERT INTO ClimbingStats Values(NULL,` + tools.QueryField(tools.GetDate()) + `,` + tools.QueryField(uData.UserID) + `, 0, 0, 0, 0, 0, 0, 0);`
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
	myquery := `INSERT INTO PullUpDB VALUES (NULL, ` + tools.QueryField(tools.GetDate()) + `, ` + tools.QueryField(uData.UserID) + `, 0, 0);`
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
