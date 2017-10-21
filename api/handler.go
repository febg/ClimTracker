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
	defer log.Printf("-> [INFO] Registration request terminated")
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
	defer log.Printf("-> [INFO] Login request terminated")

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
