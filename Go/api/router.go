package api

import "github.com/gorilla/mux"

// StandardRouter Manages request to different API endpoints
func StandardRouter(c *Control) *mux.Router {
	r := mux.NewRouter()

	r.Methods("GET").Path("/register/{user_email}/{user_password}/{user_name}").HandlerFunc(c.PostRegisterUser)
	r.Methods("GET").Path("/login/{user_email}/{user_password}").HandlerFunc(c.PostLogInUser)
	//r.Methods("POST").Path("/login/{user_emial}/{user_password}").HandlerFunc(c.PostLogInUser)
	return r
}
