package tools

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"time"

	qrcode "github.com/skip2/go-qrcode"
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

func QueryField(t string) string {
	return "'" + t + "'"
}

// Getdate returns string with current date information in 2006-01-2 format
func GetDate() string {
	return time.Now().Local().Format("2006-01-02")
}

func GenerateQrCode(fID string) string {
	log.Print("-> [TOOLS] Encoding QR code...")
	qr, err := qrcode.Encode(fID, qrcode.Medium, 256)
	if err != nil {
		log.Printf("-> [ERROR] %v", err)
		return ""
	}
	//qrcode.WriteFile("content", qrcode.Medium, 256, "qr.png")
	return base64.StdEncoding.EncodeToString(qr)
}

// func FindDateEntry(d *gym.ClimbingData, id string) bool {
// 	for _, v := range d.Data {
// 		if v.Day == id {
// 			return true
// 		}
// 	}

// 	return false
// }

func FindInSlice(in []interface{}, id string) bool {
	for _, v := range in {
		if v == id {
			return true
		}
	}

	return false
}

func StringtoInterferance(s []string) []interface{} {
	out := make([]interface{}, len(s))
	for i, v := range s {
		out[i] = v
	}
	return out
}

func NewJSON(data interface{}) ([]byte, error) {
	bs, err := json.Marshal(data)
	if err != nil {
		log.Printf("[FATAL] Unable to Marshal request: %v", err)
		return nil, err
	}
	return bs, nil
}
