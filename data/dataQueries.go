package data

import (
	"database/sql"
	"encoding/json"
	"log"
)

//ValidateRagistration executes MySQL query to check user email in database
func ValidateRagistration(DB *sql.DB, uData []byte) (bool, error) {
	//  var rows *sql.Rows
	var data UserData
	err := json.Unmarshal(uData, &data)
	if err != nil {
		log.Printf("->[ERROR] Unable to Unmarshal user information: %v", err)
		//w.WriteHeader(http.StatusBadRequest)
		//fmt.Fprint(w, "Bad Request")
		return false, err
	}

	qe := "'" + data.Email + "'"
	rows, err := DB.Query(`SELECT email FROM users WHERE email=` + qe + `;`)
	if err != nil {
		log.Printf("-> [ERROR] Registration Query: %v, %v", err, DB)
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
