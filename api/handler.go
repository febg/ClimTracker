package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"github.com/febg/Climbtracker/data"
	"reflect"

	"../data"
	"github.com/gorilla/mux"
)

// PostRegisterUser gets new user information from HTTP Post request and registers user in the main user Database
func (c *Control) PostRegisterUser(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	log.Printf("Type of marshal: %v", reflect.TypeOf(v))
	uD := data.UserData{
		Name:     v["user_name"],
		Email:    v["user_email"],
		Password: v["user_password"],
	}

	log.Printf("-> [REQUEST] Registration Request for user: %v", uD.Email)
	defer log.Printf("-> [REQUEST] User: %v successfully registered", uD.Email)

	if uD.Name == "" || uD.Email == "" || uD.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ERROR: No User Id was given")
		return
	}
	b, err := json.Marshal(uD)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, string("Internal Server Error"))
		log.Printf("[FATAL] Unable to Marshal request: %v", err)
		return
	}
	if succes, err := data.CheckUserExistance(c.DataBase, b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Internal Server error")
		return
	} else if succes != true {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: Emal already registered in data base, use another email or log in")
		return
	}

	if data.RegisterUser(c.DataBase, b) {

	}
	return
}

// PostLogInUser gets clients credentials from HTTP pPost request, compares it with existing credentials on database and grants or denies access
func (c *Control) PostLogInUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("[LOG] Login in User")
}
