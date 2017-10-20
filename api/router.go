package api

import "github.com/gorilla/mux"

// StandardRouter Manages request to different API endpoints
func StandardRouter(c *Control) *mux.Router {
	r := mux.NewRouter()

	r.Methods("GET").Path("/register/").HandlerFunc(c.PostRegisterUser)
	r.Methods("POST").Path("/login/{user_emial}/{user_password}").HandlerFunc(c.PostLogInUser)

	return r
}
