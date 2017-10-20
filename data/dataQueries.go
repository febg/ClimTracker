package data

import (
	"database/sql"
	"log"
)

//ValidateRagistration executes MySQL query to check user email in database
func ValidateRagistration(DB *sql.DB, uData UserData) (bool, error) {
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
func SendUser(DB *sql.DB, uData UserData) (bool, error) {
	myQuery := `INSERT INTO users VALUES (NULL,` + `'` + uData.Name + `','` + uData.Email + `','` + uData.Password + `', NULL);`

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
