package tools

import (
	"log"
	//"github.com/febg/Climbtracker/tools"
	"golang.org/x/crypto/bcrypt"
)

// EncryptPassword encrypts user password to store in database
func EncryptPassword(pwd string) string {
	log.Print("-> [TOOLS] Encrypting user password")
	bs, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Printf("-> [ERROR] %v", err)
		return ""
	}
	return string(bs)
}

// ComparePasswords Compares user provided password with hashed stores password
func ComparePasswords(pwd string, encPwd string) bool {
	log.Print("-> [TOOLS] Decrypting user password...")
	err := bcrypt.CompareHashAndPassword([]byte(encPwd), []byte(pwd))
	if err != nil {
		log.Printf("-> [ERROR] %v", err)
		return false
	}
	return true
}

// Boulder converts block level to database field notation
func Boulder(l string) string {
	return "V" + l
}

// QueryTable formats string containing email address to MySQL query standards
func QueryTable(t string) string {
	return "`" + t + "`"
}
