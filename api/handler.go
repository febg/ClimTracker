package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// PostRegisterUser gets new user information from HTTP Post request and registers user in the main user Database
func (c *Control) PostRegisterUser(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request")
		return
	}
	log.Printf("[LOG] Registering User")

	var action string
	err = json.Unmarshal(reqBytes, &action)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request")
		return
	}
}

// PostLogInUser gets clients credentials from HTTP pPost request, compares it with existing credentials on database and grants or denies access
func (c *Control) PostLogInUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("[LOG] Login in User")
}
