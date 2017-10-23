package data

import (
	"database/sql"
	"log"

	"github.com/febg/Climbtracker/Go/tools"
)

//ValidateRagistration executes MySQL query to check user email in database
func validateRagistration(DB *sql.DB, uData UserData) (bool, error) {
	qe := "'" + uData.Email + "'"
	rows, err := DB.Query(`SELECT email FROM users WHERE email=` + qe + `;`)
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
	myQuery := `INSERT INTO users VALUES (NULL,` + `'` + uData.Name + `','` + uData.Email + `','` + uData.Password + `','` + uData.UserID + `');`

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

// CreateUserTable executes query to create a a table for newly registered user
func createUserTable(DB *sql.DB, uData UserData) (bool, error) {
	table := "`" + uData.UserID + "`"
	myQuery := `CREATE TABLE ` + table + ` (uid INT NOT NULL UNIQUE AUTO_INCREMENT, date VARCHAR(20) NOT NULL UNIQUE, V1 INT, V2 INT, V3 INT, V4 INT, V5 INT, V6 INT, PRIMARY KEY (uid));`
	stmt, err := DB.Prepare(myQuery)
	if err != nil {
		log.Printf("-> [ERROR] Create Table query preparation: %v", err)
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Printf("-> [ERROR] Create Table query execution: %v", err)
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

func getClimbingData(DB *sql.DB, uID string) {
	qe := "`" + "felipeb85@gmail.com" + "`"
	rows, err := DB.Query(`SELECT * FROM ` + qe + `;`)
	if err != nil {
		log.Printf("-> [ERROR] Get user climbing data query: %v", err)
		return
	}
	defer rows.Close()
	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		log.Printf("-> [ERROR] Get user climbing data query: %v", err)
		return // proper error handling instead of panic in your app
	}
	values := make([]sql.RawBytes, len(columns))
	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

}

func recordBlock(DB *sql.DB, cData NewCheckIn) error {
	stmt, err := DB.Prepare(`UPDATE ` + tools.QueryTable(cData.UserID) + ` SET ` + tools.Boulder(cData.Level) + `=` + tools.Boulder(cData.Level) + `+ 1 WHERE date=` + tools.GetDate() + `;`)
	if err != nil {
		log.Printf("-> [ERROR] Record Block query preparation: %v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Printf("-> [ERROR] Record Block query execution: %v", err)
	}
	return nil
}

func initializeTable(DB *sql.DB, bData NewCheckIn) error {
	date := "'" + tools.GetDate() + "'"
	log.Printf(date)
	myquery := `INSERT IGNORE INTO ` + tools.QueryTable(bData.UserID) + ` SET date = ` + date + `, V1 = 0, V2 = 0, V3 = 0, V4 = 0, V5 = 0, V6 = 0;`
	stmt, err := DB.Prepare(myquery)
	if err != nil {
		log.Printf("-> [ERROR] Initialize Table query preparation: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Printf("-> [ERROR] Initialize Table query execution: %v", err)
		return err
	}
	return nil
}
