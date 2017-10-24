package api

import "github.com/gorilla/mux"

// StandardRouter Manages request to different API endpoints
func StandardRouter(c *Control) *mux.Router {
	r := mux.NewRouter()

	r.Methods("GET").Path("/register/{user_email}/{user_password}/{user_name}/").HandlerFunc(c.PostRegisterUser)
	r.Methods("POST").Path("/login/{user_email}/{user_password}/").HandlerFunc(c.PostLogInUser)
	r.Methods("GET").Path("/checkin/{user_id}/{level}").HandlerFunc(c.PostCheckIn)
	r.Methods("GET").Path("/getall/{user_id}/").HandlerFunc(c.PostGetData)
	r.Methods("GET").Path("/getfriends/{user_id}/").HandlerFunc(c.PostGetData)
	r.Methods("GET").Path("/addfriend/{user_id}/{friend_email}/{qr}").HandlerFunc(c.PostGetData)
	r.Methods("GET").Path("/recordpullup/{user_id}/{amount}").HandlerFunc(c.PostGetData)
	return r
}
