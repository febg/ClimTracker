package data

import (
	"database/sql"
	"log"
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
	myQuery := `CREATE TABLE ` + table + ` (uid INT NOT NULL UNIQUE AUTO_INCREMENT, date VARCHAR(20) NOT NULL, V1 INT, V2 INT, V3 INT, V4 INT, V5 INT, V6 INT, PRIMARY KEY (uid));`
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
