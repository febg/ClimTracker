package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/febg/Climbtracker/data"
	"github.com/febg/Climbtracker/tools"
	//"../data"
	//"../tools"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

// PostRegisterUser gets new user information from HTTP Post request and registers user in the main user Database
func (c *Control) PostRegisterUser(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	log.Printf("[REQUEST] Registration request for user: %v", v["user_name"])
	defer log.Printf("----------------------------------------")
	defer log.Printf("[REQUEST] Registration request terminated")
	uD := data.UserData{
		Name:     v["user_name"],
		Email:    v["user_email"],
		Password: tools.EncryptPassword(v["user_password"]),
		UserID:   uuid.NewV4().String(),
	}

	if uD.Name == "" || uD.Email == "" || uD.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ERROR: Registration information given not complete")
		return
	}
	b, err := json.Marshal(uD)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, string("Internal Server Error"))
		log.Printf("[FATAL] Unable to Marshal request: %v", err)
		return
	}

	if newUser, err := data.NewUser(c.DataBase, b); newUser != true {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Internal Server error")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: Email already registered in data base, use another email or log in")
		return
	}
	fmt.Fprintf(w, "Success: UserID: %v", uD.UserID)
	return
}

func (c *Control) PostGetData(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	uID := v["user_id"]
	log.Printf("[REQUEST] Data request for user: %v", uID)
	defer log.Printf("----------------------------------------")
	defer log.Printf("[REQUEST] Data request terminated")

	if uID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("-> [ERROR] Check in information not complete")
		fmt.Fprint(w, string("Internal Server Error"))
		return
	}
	log.Printf("-> [INFO] Getting user's climbing history..")
	cData, err := data.ClimbingHistory(c.DataBase, uID)
	if err != nil {
		fmt.Fprint(w, string("Internal Server Error"))
		log.Printf("[ERROR] Unable to obtain Climbhistory: %v", err)
	}
	log.Printf("-> [INFO] Climbing history successfully obtained")
	log.Printf("-> [INFO] Encoding climbing history...")
	b, err := json.Marshal(cData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, string("Internal Server Error"))
		log.Printf("[FATAL] Unable to Marshal request: %v", err)
		return
	}
	log.Printf("-> [INFO] Climbing history sent to client successfully")
	fmt.Fprint(w, string(b))

	return
}

// PostLogInUser gets clients credentials from HTTP pPost request, compares it with existing credentials on database and grants or denies access
func (c *Control) PostLogInUser(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	uD := data.UserData{
		Name:     "",
		Email:    v["user_email"],
		Password: v["user_password"],
		UserID:   "",
	}

	log.Printf("[REQUEST] Login request for user: %v", uD.Email)
	defer log.Printf("----------------------------------------")
	defer log.Printf("[REQUEST] Login request terminated")

	if uD.Email == "" || uD.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("-> [ERROR] Log in information not complete")
		fmt.Fprint(w, "ERROR: Log in information not complete")
		return
	}
	b, err := json.Marshal(uD)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, string("Internal Server Error"))
		log.Printf("[FATAL] Unable to Marshal request: %v", err)
		return
	}
	uID, err := data.LogIn(c.DataBase, b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, string("Internal Server Error"))
		return
	}
	if uID == "wpwd" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ERROR: Wrong user name or password")
		return
	}
	fmt.Fprintf(w, "SUCCESS: UserID: %v", uID)
	return
}

// PostCheckIn handles request to store a climbing block
func (c *Control) PostCheckIn(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	uD := data.NewCheckIn{
		Level:  v["level"],
		UserID: v["user_id"],
	}

	log.Printf("[REQUEST] Check in request for user: %v", uD.UserID)
	defer log.Printf("----------------------------------------")
	defer log.Printf("[REQUEST] Check in request terminated")

	if uD.Level == "" || uD.UserID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("-> [ERROR] Check in information not complete")
		fmt.Fprint(w, "ERROR: Check in information not complete")
		return
	}
	bs, err := json.Marshal(uD)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, string("Internal Server Error"))
		log.Printf("[FATAL] Unable to Marshal request: %v", err)
		return
	}
	data.CheckIn(c.DataBase, bs)
	//Handle errors and respond to client
}
