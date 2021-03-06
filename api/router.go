package api

import "github.com/gorilla/mux"

// StandardRouter Manages request to different API endpoints
func StandardRouter(c *Control) *mux.Router {
	r := mux.NewRouter()
	//TODO Change methods to POST when exposed to AWS mysql and servers
	r.Methods("GET").Path("/register/{user_email}/{user_password}/{user_name}/").HandlerFunc(c.PostRegisterUser)
	r.Methods("GET").Path("/login/{user_email}/{user_password}/").HandlerFunc(c.PostLogInUser)
	r.Methods("GET").Path("/checkin/{user_id}/{level}").HandlerFunc(c.PostCheckIn)
	r.Methods("GET").Path("/getall/{user_id}/").HandlerFunc(c.PostGetData)
	r.Methods("GET").Path("/getfriends/{user_id}/").HandlerFunc(c.PostGetFriends)
	r.Methods("GET").Path("/addfriend/{user_id}/{user_email}/").HandlerFunc(c.PostAddFriend)
	r.Methods("GET").Path("/recordpullup/{user_id}/{count}").HandlerFunc(c.PostRecordPullUp)
	return r
}
